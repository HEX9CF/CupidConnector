# 🔗 Cupid Connector

## 项目简介

汕头大学校园网自动登录与流量监控第三方轻量级工具，通过发送登录请求并处理响应，实现自动化登录功能，方便汕大师生登录校园网。同时提供一个流量信息仪表面板，支持流量监控、流量告警、自动注销等功能，防止流量超额使用。将程序添加到开机启动项，即可实现开机自动登录校园网。

A lightweight third-party tool for automatically logging into the Shantou University campus network. It automates the login process by sending login requests and handling responses, making it convenient for Shantou University students and staff to log into the campus network. It provides a traffic information dashboard and supports traffic monitoring, traffic alerts, and automatic logout to prevent excessive traffic usage. Adding the program to the startup items enables automatic login to the campus network upon startup.

## 软件版本

### Cupid Connector

完整 GUI 程序，支持流量监控、流量告警、自动注销等功能

仓库地址：[https://github.com/HEX9CF/CupidConnector](https://github.com/HEX9CF/CupidConnector)

### Cupid Connector Lite

轻量版命令行工具，仅提供自动登录功能，可配合自动退出和开机自启使用

仓库地址：[https://github.com/HEX9CF/CupidConnectorLite](https://github.com/HEX9CF/CupidConnectorLite)

## 功能简介

- **自动登录**：通过发送登录请求，并处理登录响应，实现自动化登录。
- **重新认证**：用户可以在需要时重新进行登录认证。
- **注销账号**：用户可以注销当前登录的账号，退出校园网。
- **自动退出**：用户可以设置登录成功后自动退出程序。
- **自动隐藏**：用户可以设置启动程序后自动最小化窗口到系统托盘。
- **仪表面板**：提供一个流量信息仪表面板，显示当前账号信息和网络流量使用情况。
- **流量监控**：用户可以开启流量监控功能，设置监控间隔时间，监控网络流量使用情况。
- **断线重连**: 当流量监控检测到校园网连接断开时，自动尝试重新登录校园网，保证你的文件上传下载和远程桌面连接持续在线。
- **流量告警**：用户可以设置流量告警阈值，当网络流量超过阈值时，会弹出告警提示，提醒用户注意流量使用。
- **自动注销**：用户可以设置自动注销阈值，当网络流量超过阈值时，自动注销账号，防止超额使用流量。

## 使用方法

1. **运行程序**：在 Release 页面下载最新的可执行文件，双击运行程序。
2. **设置账号**：打开设置页面，输入校园网账号和密码，点击保存。
3. **自动登录**：程序会自动发送认证请求，登录校园网，登录成功后会显示登录成功的提示信息。
4. **自动退出**：用户可以设置登录成功后自动退出程序，也可以手动点击退出按钮退出程序。若需取消自动退出，可以将配置文件中的 `AUTO_EXIT` 设置为 `FALSE`。
5. **流量监控**：用户可以开启流量监控功能，设置监控间隔时间和告警阈值，监控网络流量使用情况。
6. **开机自启**：用户可以将程序添加到开机启动项，实现开机自动登录校园网。具体方法请参考下文。

## 配置文件

配置文件为 `stu.env`，位于程序运行目录下。 配置文件内容如下：

```env
# 校园网认证接口
BASE_URL=https://a.stu.edu.cn

# 校园网用户名
STU_USERNAME=username
# 校园网密码
STU_PASSWORD=password

# 启动程序后自动登录
AUTO_LOGIN=TRUE
# 认证成功后自动退出程序
AUTO_EXIT=FALSE
# 启动程序后自动隐藏窗口
AUTO_HIDE=FALSE
# 断线自动重连校园网
AUTO_RECONNECT=TRUE

# 自动监控网络流量
MONITOR_FLUX=TRUE
# 监控间隔时间（分钟）
MONITOR_INTERVAL=5
# 监控告警阈值（剩余流量百分比）
MONITOR_ALERT_THRESHOLD=30
# 自动注销阈值（剩余流量百分比）
MONITOR_LOGOUT_THRESHOLD=10
```

## 开机自启

用户可以将程序添加到开机启动项，实现开机自动登录校园网。

这里以 Windows 系统为例，具体步骤如下：

1. 将可执行文件移动到一个固定的目录，如 `C:\Program Files\Cupid Connector`。
2. 右键点击可执行文件 `cupid-connector.exe`，选择“创建快捷方式”。
3. 按 `Win + R` 打开运行对话框。
4. 输入 `shell:startup` 并按回车，打开启动文件夹。 
5. 将刚才创建的快捷方式移动到这个文件夹中。

如果需要取消开机自启，只需将快捷方式移出启动文件夹即可。

## 杀毒误报

本项目是开源项目，用户可以自行审查源代码，确认程序的安全性。用户的账号密码等信息只会保存在用户的计算机上，仅用于校园网登录认证。

由于程序需要发送网络请求，可能会被杀毒软件误报为恶意软件。如果遇到杀毒软件误报，请将程序添加到信任列表中。

## 软件界面

### 主界面
![image](https://github.com/user-attachments/assets/a683aa37-bd7a-4b55-bd03-77dec098376b)

### 基本设置
![image](https://github.com/user-attachments/assets/2e35cf80-4820-49b6-8907-4070e34263aa)

### 流量监控
![image](https://github.com/user-attachments/assets/fb500d22-2a25-48f2-b2de-8d22352b61bc)

### 流量告警通知框
![image](https://github.com/user-attachments/assets/f58eb0f7-9f70-4af4-a05a-6a6ea4726a67)

## 趣闻

项目名称 "Cupid Connector" 灵感来源于游戏《Death Stranding》中的 "丘比特连接器"（Q-pid）。在游戏中，丘比特连接器是一种设备，用于连接和扩展开罗尔网络，使得玩家能够与更广阔的世界进行通信和互动。这个设备在游戏中扮演着连接和桥梁的角色，象征着连接和沟通的重要性。Cupid Connector 作为一个自动登录校园网的工具，同样扮演着连接的角色。它将学生和教职工与校园网络连接起来，使他们能够方便地访问网络资源，寓意它能够像游戏中的丘比特连接器一样，为用户架起一座连接的桥梁。

## 参与贡献

- 如果你觉得这个项目对你有所帮助，欢迎给个 Star⭐️。
- 如果你有任何问题或建议，欢迎提交 Issue。
- 如果你有兴趣参与贡献代码，欢迎提交 Pull Request。

## 许可与声明

本项目采用 GPL-3.0 license 进行许可，详情请参阅 [LICENSE](LICENSE) 文件。

本项目仅供学习交流使用，不得用于任何商业用途。使用本项目所造成的任何后果，由使用者自行承担，与项目作者无关。
