通过Bazel用来管理所有学习的monorepo工程

## Bazel
bazel run //:gazelle -- update-repos -from_file=go.mod

已有项目通过git submodule的方式来引入
- ical4notion
- auto-dev-rag-demo
- elonws

## Submodule 如何提交代码
- 通过git submodule foreach git pull 来更新子模块
- 通过git submodule foreach git add . 来添加子模块的修改
- 通过git submodule foreach git commit -m "message" 来提交子模块的修改
- 通过git submodule foreach git push 来推送子模块的修改
- 通过git push origin master 来推送父模块的修改
- git push --recurse-submodules=on-demand 来推送父模块和子模块的修改
- 通过git submodule update --remote --merge 来更新子模块的远程分支


## gliese.cn 链路
d.gliese.cn -(cloudflare)-> Tunnels -(cloudflare)-> ubuntu@server:8080 --> localhost:8080(fpr)

plex.gliese.cn -(cloudflare)-> Tunnels -(cloudflare)-> ubuntu@SERVER_IP:32400 -> localhost:32400

s.gliese.cn -(cloudflare)-> Tunnels -(cloudflare)-> ubuntu@SERVER_IP:80

原则：本地尽量少部署服务，尽量使用云服务

## ChangeLog
- 2021-09-06: 项目初始化
- 2024-06-22: 使用Github Copilot来

