# 汕头大学校园网自动化登录工具

## 项目简介

汕头大学校园网自动化登录工具是一个用 Go 语言编写的软件，该工具通过模拟登录请求，并处理登录响应，从而实现自动化登录功能。

Shantou University Campus Network Auto Login Tool is a software written in Go language. This tool automates the login process by simulating login requests and handling login responses.

## 功能简介

- **自动登录**：通过发送登录请求，并处理登录响应，实现自动化登录。
- **重新认证**：用户可以在需要时重新进行登录认证。
- **注销账号**：用户可以注销当前登录的账号，退出校园网。
- **初始化配置**：支持初始化配置文件，简化用户的配置过程。

## 使用方法

1. **运行程序**：下载本项目的可执行文件，双击运行程序。
2. **初始化配置**：首次运行程序时，会自动生成配置文件。用户可以根据提示填写校园网用户名和密码。同时，程序会询问是否在认证成功后自动退出程序，用户可以根据需求选择是否开启该功能。
3. **自动登录**：程序会自动发送认证请求，登录校园网，登录成功后会显示登录成功的提示信息。
4. **自动退出**：如果用户开启了认证成功自动退出功能，程序会在认证成功后立即自动退出，不会进入命令模式。如果没有认证成功，仍会进入命令模式。
5. **命令模式**: 如果用户没有开启认证成功自动退出功能，在自动登录完成后，程序会进入命令模式，用户可以输入命令来执行相应的操作。
6. **开机自启**：用户可以将程序添加到开机启动项，实现开机自动登录校园网。

## 命令列表

- `login`：登录账号。
- `logout`：注销账号。
- `init`: 重新初始化配置文件。
- `exit`：退出程序。

## 配置文件

配置文件为 `stu.env`，位于程序运行目录下。 配置文件内容如下：

```env
# 校园网认证接口
STU_URL=http://a.stu.edu.cn/ac_portal/login.php

# 校园网用户名
STU_USERNAME=username

# 校园网密码
STU_PASSWORD=password

# 认证成功后自动退出程序
AUTO_EXIT=FALSE
```
