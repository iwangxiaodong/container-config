
构建包括扩展的Fn Server镜像：

    wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/myext/myext.go
    
    vim ext.yaml

```yaml
extensions:
  - name: github.com/iwangxiaodong/container-config/master/fn/myext
```
  
    fn build-server -t cr.openle.com/my/fn-server-with-ext
