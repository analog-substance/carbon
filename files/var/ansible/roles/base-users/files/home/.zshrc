export ZSH=$HOME/.oh-my-zsh
export NMAP_PRIVILEGED=""

ZSH_THEME="robbyrussell"
plugins=(git wd)

source $ZSH/oh-my-zsh.sh

ls  ~/.zshrc.d/*.zshrc | sort -n | while read f; do
  source "$f";
done
