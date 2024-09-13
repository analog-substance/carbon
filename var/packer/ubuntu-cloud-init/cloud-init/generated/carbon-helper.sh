#! /bin/bash
TMP_DIR_PREFIX="/run/carbon"

function temp_dir_init () {
    mkdir -p $TMP_DIR_PREFIX/$1
    cd $TMP_DIR_PREFIX/$1
}

function temp_dir_clean () {
    cd - 2>&1 > /dev/null
    rm -rf $TMP_DIR_PREFIX/$1
}


function install_ohmyzsh () {
    if [ ! -d /etc/skel/.oh-my-zsh ]; then
        git clone -c core.autocrlf=input --depth=1 https://github.com/robbyrussell/oh-my-zsh.git /etc/skel/.oh-my-zsh
    fi
}

function nmap_privileges () {
    setcap cap_net_raw,cap_net_admin,cap_net_bind_service+eip /usr/bin/nmap
    echo "export NMAP_PRIVILEGED=1" > /etc/skel/.zshrc.d/nmap_privs.zshrc
}

function bootstrap_user () {
    rsync -a /etc/skel/ /home/ubuntu --exclude .vnc
    chown -R ubuntu:ubuntu /home/ubuntu
    chsh -s /usr/bin/zsh ubuntu
}

function install_go () {
    temp_dir_init $FUNCNAME

    go_url="https://golang.org/dl/$(curl -s 'https://go.dev/dl/?mode=json' | jq -r '.[0].version').linux-amd64.tar.gz"
    curl -o go.linux-amd64.tar.gz -Ls "$go_url"
    rm -rf /usr/local/go
    tar -C /usr/local -xzf go.linux-amd64.tar.gz

    temp_dir_clean $FUNCNAME
}

function setup_opt () {
    mkdir -p /opt/go 
    chown -R ubuntu:sudo /opt 
    chmod -R 775 /opt 
    setfacl -R -m g:sudo:rwX /opt; find /opt -type d -exec setfacl -R -m d:g:sudo:rwX "{}" \;
}

function clone_repo () {
    git clone --depth 1 $1
}

function install_arsenic () {
    git clone --depth 1 https://github.com/analog-substance/arsenic.git /opt/arsenic
    cd /opt/arsenic
    GOCACHE=/opt/go/.cache GOPATH=/opt/go /usr/local/go/bin/go install
    cd -
}

function go_install () {
    name=$(basename $1 | cut -d '@' -f1)
    GOCACHE=/opt/go/.cache GOPATH=/opt/go /usr/local/go/bin/go install -v $1 2>&1 | sed 's/^/['"$name"'] /'
}

function install_github_release () {
    temp_dir_init $FUNCNAME
    set +e
    repo=$1
    search=$2

    os='linux'
    initial_filter=".[0].assets[]"
    initial_filter="$initial_filter | select((.name | ascii_downcase | contains(\"$os\")))"

    if [[ "$(uname -m)" == "x86_64" ]]; then
        initial_filter="$initial_filter | select((.name | ascii_downcase | contains(\"amd64\")))" # or (.name | ascii_downcase | contains(\"64bit\")))"
    fi

    if [ ! -z "$search" ]; then
        echo "appending filter for $search"
        initial_filter="$initial_filter | select((.name | ascii_downcase | contains(\"$search\")))"
    fi

    res=$(curl -s "https://api.github.com/repos/$repo/releases" | jq -r "$initial_filter | .browser_download_url")
    zip=$(echo "$res" | grep \.zip)
    tar=$(echo "$res" | grep -P "\.tar\.gz|\.tgz")
    deb=$(echo "$res" | grep \.deb)

    mkdir -p "$repo"
    cd "$repo"
    set -e

    if [ ! -z "$zip" ]; then 
	echo "installing zip ($zip)"
        curl -o install.zip -Ls "$zip"
        unzip install.zip

        find . -type f -executable -exec mv {} /usr/local/bin/ \;
    elif [ ! -z "$tar" ]; then 
	echo "installing tar ($tar)"
        curl -o install.tgz -Ls "$tar"
        tar -xzvf install.tgz

        find . -type f -executable -exec mv {} /usr/local/bin/ \;
    elif [ ! -z "$tar" ]; then 
	echo "installing deb ($deb)"
        curl -o install.deb -Ls "$deb"
        dpkg -i install.deb
    fi
    
    cd -
    temp_dir_clean $FUNCNAME
}

function install_go_tools () {
    for tool in $(cat /etc/carbon/go-installs); do
        go_install "$tool"
    done
}

function install_github_releases () {
    cat /etc/carbon/github-releases | while read line; do
        tool=$(echo $line | awk '{print $1}')
        search=$(echo $line | awk '{print $2}')
        install_github_release "$tool" "$search"
    done
}

function install_git_repos () {
    cd /opt
    for repo in $(cat /etc/carbon/git-repos); do
        clone_repo "$repo"
    done
    cd -
}

function install_aws_cli () {
    wget https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip 
    unzip awscli-exe-linux-x86_64.zip > /dev/null
    chmod +x aws/install 
    aws/install 
    rm -rf aws*
}

function clean_up () {
    rm -rf $TMP_DIR_PREFIX
    rm -rf /opt/go/.cache
}

function init_user () {
    install_ohmyzsh
    nmap_privileges
    bootstrap_user
}

function main () {
    set -e
    set -x
    install_go
    setup_opt

    init_user
    install_aws_cli
    install_arsenic

    install_go_tools
    install_github_releases
    install_git_repos
    clean_up
}

time main
