---
title: ssh_util
description: 
weight: 200
---


```go
import "github.com/analog-substance/carbon/pkg/ssh_util"
```

## Index

- [type Session](<#Session>)
  - [func NewSession\(\) \(\*Session, error\)](<#NewSession>)
  - [func \(session \*Session\) ClientConfig\(user string\) \*ssh.ClientConfig](<#Session.ClientConfig>)
  - [func \(session \*Session\) Close\(\)](<#Session.Close>)
  - [func \(session \*Session\) Connect\(serverAddr, user string\) error](<#Session.Connect>)
  - [func \(session \*Session\) ForwardAgent\(\) error](<#Session.ForwardAgent>)
  - [func \(session \*Session\) ForwardLocalPort\(localPort, remotePort int\) error](<#Session.ForwardLocalPort>)
  - [func \(session \*Session\) Output\(cmd string\) \(string, error\)](<#Session.Output>)


<a name="Session"></a>
## type [Session](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L17-L22>)



```go
type Session struct {
    Session *ssh.Session
    Client  *ssh.Client
    // contains filtered or unexported fields
}
```

<a name="NewSession"></a>
### func [NewSession](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L24>)

```go
func NewSession() (*Session, error)
```



<a name="Session.ClientConfig"></a>
### func \(\*Session\) [ClientConfig](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L54>)

```go
func (session *Session) ClientConfig(user string) *ssh.ClientConfig
```



<a name="Session.Close"></a>
### func \(\*Session\) [Close](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L36>)

```go
func (session *Session) Close()
```



<a name="Session.Connect"></a>
### func \(\*Session\) [Connect](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L66>)

```go
func (session *Session) Connect(serverAddr, user string) error
```



<a name="Session.ForwardAgent"></a>
### func \(\*Session\) [ForwardAgent](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L83>)

```go
func (session *Session) ForwardAgent() error
```



<a name="Session.ForwardLocalPort"></a>
### func \(\*Session\) [ForwardLocalPort](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L96>)

```go
func (session *Session) ForwardLocalPort(localPort, remotePort int) error
```



<a name="Session.Output"></a>
### func \(\*Session\) [Output](<https://github.com/analog-substance/carbon/blob/main/pkg/ssh_util/main.go#L143>)

```go
func (session *Session) Output(cmd string) (string, error)
```

Output uses ssh\_util.Session to run cmd on the remote host and returns its standard output.

