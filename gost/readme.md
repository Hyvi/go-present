# Gost SOCKS5+TLS 代理部署指南 (Ubuntu)

本文档记录了在 Ubuntu 服务器上部署 Gost 代理服务的完整步骤，配置为 `socks5+tls` 协议，端口 `7500`，并开启 UDP 支持。

## 1. 安装 Gost (v2.12)

snap install core gost 
# 4. 验证版本 (应显示 v2.12)
gost -V
```

## 2. 配置 Systemd 服务

使用 Systemd 管理服务，确保开机自启和后台运行。

创建或覆盖 `/etc/systemd/system/gost.service` 文件：

```bash
sudo bash -c 'cat > /etc/systemd/system/gost.service <<EOF
[Unit]
Description=Gost Proxy Service
After=network.target

[Service]
Type=simple
User=root
# 配置说明:
# 协议: socks5+tls (SOCKS5 over TLS)
# 端口: 7500
# 用户: admin
# 密码: <密码>
# 参数: udp=true (开启 UDP 转发支持)
ExecStart=/usr/bin/gost -L "socks5+tls://admin:密码@:7500?udp=true"
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF'
```

## 3. 启动服务

```bash
# 1. 重载 Systemd 配置
sudo systemctl daemon-reload

# 2. 启动服务并设置为开机自启
sudo systemctl enable --now gost

# 3. 检查运行状态
sudo systemctl status gost
```

## 4. 验证部署

### 检查端口监听
确认 7500 端口是否被 `gost` 进程监听：
```bash
sudo ss -tlnp | grep :7500
```
*输出应包含 `LISTEN ... :7500 ... users:(("gost",...))`*

### 检查日志
查看最近的运行日志，确认没有报错：
```bash
sudo journalctl -u gost -n 20 --no-pager
```
*正常日志应显示 `socks5+tls://:7500 on [::]:7500`*

## 5. 客户端连接配置

在 Shadowrocket / v2rayNG / Clash 等客户端中添加如下配置：

*   **类型 (Type)**: `SOCKS5` (或 `SOCKS5 over TLS`)
*   **地址 (Address)**: `服务器公网IP`
*   **端口 (Port)**: `7500`
*   **用户 (User)**: `admin`
*   **密码 (Password)**: `密码`
*   **TLS**: **开启 (Enable)**
*   **跳过证书验证 (Allow Insecure)**: **开启 (True)** *(因为使用的是自签名证书)*
*   **UDP 转发**: 开启 (如果客户端支持)

## 6. 注意事项

1.  **防火墙/安全组**: 务必在云服务商（AWS/阿里云/腾讯云等）的控制台安全组中，放行 **TCP 7500** 和 **UDP 7500** 端口。
2.  **安全性**: 当前密码已硬编码在服务文件中，如需修改，请编辑 `/etc/systemd/system/gost.service` 并运行 `sudo systemctl daemon-reload && sudo systemctl restart gost`。
