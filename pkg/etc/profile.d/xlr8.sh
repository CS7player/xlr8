xlr8_go() {
    xlr8

    if [ -f /tmp/xlr8-cwd ]; then
        cd "$(cat /tmp/xlr8-cwd)"
    fi
}
