http-over-ssh-example
---------------------

Usage
=====

    go build
    ./http-over-ssh-example -ssh.user=$USERNAME -ssh.pass=$PASSWORD
    ab -c 100 -n 10000 http://localhost:8080/ 
