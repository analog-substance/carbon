idempotent_path_prepend ()
{
    PATH=${PATH//":$1"/} #delete any instances in the middle or at the end
    PATH=${PATH//"$1:"/} #delete any instances at the beginning
    export PATH="$1:$PATH" #prepend
}

idempotent_path_prepend /usr/local/go/bin
idempotent_path_prepend /opt/go/bin

export GOPATH="/opt/go"
