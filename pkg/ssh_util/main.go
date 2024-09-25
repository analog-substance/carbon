package ssh_util

import (
	"fmt"
	sshagent "github.com/xanzy/ssh-agent"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// Thanks https://eli.thegreenplace.net/2022/ssh-port-forwarding-with-go/

type Session struct {
	Session         *ssh.Session
	Client          *ssh.Client
	agent           agent.Agent
	agentConnection net.Conn
}

func NewSession() (*Session, error) {
	newAgent, conn, err := sshagent.New()
	if err != nil {
		return nil, fmt.Errorf("failed to open SSH_AUTH_SOCK: %v", err)
	}

	return &Session{
		agent:           newAgent,
		agentConnection: conn,
	}, nil
}

func (session *Session) Close() {
	if session.Client != nil {
		if err := session.Client.Close(); err != nil {
			log.Printf("error encountered while closing SSH client: %v", err)
		}
	}

	if session.Session != nil {
		if err := session.Session.Close(); err != nil {
			log.Printf("error encountered while closing SSH session: %v", err)
		}
	}

	if err := session.agentConnection.Close(); err != nil {
		log.Printf("error encountered while closing SSH agent connection: %v\n", err)
	}
}

func (session *Session) ClientConfig(user string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use a callback rather than PublicKeys, so we only consult the
			// agent once the remote server wants it.
			ssh.PublicKeysCallback(session.agent.Signers),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func (session *Session) Connect(serverAddr, user string) error {
	sshClient, err := ssh.Dial("tcp", serverAddr, session.ClientConfig(user))
	if err != nil {
		return fmt.Errorf("error dialing ssh (%s): %v", serverAddr, err)
	}

	session.Client = sshClient
	sshSession, err := session.Client.NewSession()
	if err != nil {
		return fmt.Errorf("error creating ssh session: %v", err)
	}
	sshSession.Stderr = os.Stderr
	session.Session = sshSession

	return nil
}

func (session *Session) ForwardAgent() error {
	err := agent.RequestAgentForwarding(session.Session)
	if err != nil {
		return fmt.Errorf("error requesting agent forwarding: %v", err)
	}

	err = agent.ForwardToAgent(session.Client, session.agent)
	if err != nil {
		return fmt.Errorf("error forwarding to agent: %v", err)
	}
	return nil
}

func (session *Session) ForwardLocalPort(localPort, remotePort int) error {

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", localPort))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// Like ssh -L by default, local connections are handled one at a time.
		// While one local connection is active in runTunnel, others will be stuck
		// dialing, waiting for this Accept.
		local, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Issue a dial to the remote server on our SSH client; here "localhost"
		// refers to the remote server.
		remote, err := session.Client.Dial("tcp", fmt.Sprintf("localhost:%d", remotePort))
		if err != nil {
			log.Fatal(err)
		}

		go runTunnel(local, remote)
	}
}

func runTunnel(local, remote net.Conn) {
	defer local.Close()
	defer remote.Close()
	done := make(chan struct{}, 2)

	go func() {
		io.Copy(local, remote)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(remote, local)
		done <- struct{}{}
	}()

	<-done
}

// Output uses ssh_util.Session to run cmd on the remote host and returns its standard output.
func (session *Session) Output(cmd string) (string, error) {
	output, err := session.Session.Output(cmd)
	if err == nil {
		return strings.TrimSpace(string(output)), nil
	}

	return "", err
}
