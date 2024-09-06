
if [ -d /usr/local/go/bin ]; then
  export PATH="/usr/local/go/bin:$PATH"
fi

if [ -d $HOME/go/bin ]; then
  export PATH="$HOME/go/bin:$PATH"
fi

if [ -d "$HOME/.local/bin" ]; then
  export PATH="$HOME/.local/bin:$PATH"
fi

if [ -d "$HOME/bin" ]; then
  export PATH="$HOME/bin:$PATH"
fi

if test "$SSH_AUTH_SOCK" ; then
  if [ "$SSH_AUTH_SOCK" != "$HOME/.ssh/ssh_auth_sock" ]; then
    export SSH_AUTH_SOCK="$HOME/.ssh/ssh_auth_sock"
  fi
fi

if [ -f $HOME/opt/arsenic/arsenic.rc ]; then
  source $HOME/opt/arsenic/arsenic.rc
fi

if which arsenic > /dev/null 2>&1 ; then
  if [ ! -f "${fpath[1]}/_arsenic" ]; then
    arsenic completion zsh >  "${fpath[1]}/_arsenic"
  fi
fi
autoload -U compinit; compinit
