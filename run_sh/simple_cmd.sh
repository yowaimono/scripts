#!/bin/bash

# 定义别名和函数
aliases=(
    "alias ll='ls -la'"
    "alias gs='git status'"
    "alias gc='git commit -m'"
    "alias gp='git push'"
    "alias dps='docker ps'"
    "alias dpsa='docker ps -a'"
    "alias dim='docker images'"
    "alias dsp='docker system prune -f'"
    "alias dvp='docker volume prune -f'"
    "alias dip='docker image prune -f'"
    "alias cls='clear'"
)

# 定义函数
functions=(
    "psg() { ps -ef | grep \"\$1\"; }"
    "drm() { docker stop \"\$@\"; docker rm \"\$@\"; }"
    "drmi() { docker rmi \"\$@\"; }"
    "drun() { docker run -it --rm \"\$@\"; }"
    "dexec() { docker exec -it \"\$1\" /bin/bash; }"
    "dlogs() { docker logs -f \"\$@\"; }"
)

# 将别名和函数添加到 ~/.bashrc 文件中
for alias_cmd in "${aliases[@]}"; do
    if ! grep -q "$alias_cmd" ~/.bashrc; then
        echo "$alias_cmd" >> ~/.bashrc
        echo "Added alias: $alias_cmd"
    else
        echo "Alias already exists: $alias_cmd"
    fi
done

for func_cmd in "${functions[@]}"; do
    if ! grep -q "$func_cmd" ~/.bashrc; then
        echo "$func_cmd" >> ~/.bashrc
        echo "Added function: $func_cmd"
    else
        echo "Function already exists: $func_cmd"
    fi
done

# 使别名和函数在当前 shell 中生效
source ~/.bashrc

echo "All aliases and functions have been added and are now active."

# 提示用户重新加载当前 shell
echo "Please open a new terminal or run 'source ~/.bashrc' to ensure the changes take effect."
