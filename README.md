通过Bazel用来管理所有学习的monorepo工程

已有项目通过git submodule的方式来引入
- ical4notion


## Submodule 如何提交代码
- 通过git submodule foreach git pull 来更新子模块
- 通过git submodule foreach git add . 来添加子模块的修改
- 通过git submodule foreach git commit -m "message" 来提交子模块的修改
- 通过git submodule foreach git push 来推送子模块的修改
- 通过git push origin master 来推送父模块的修改
- git push --recurse-submodules=on-demand 来推送父模块和子模块的修改
- 通过git submodule update --remote --merge 来更新子模块的远程分支

