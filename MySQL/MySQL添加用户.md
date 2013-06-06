//创建无密码的MySQL用户
CREATE USER '用户名'@'localhost'

//创建明文密码的MySQL用户
CREATE USER '用户名'@'localhost' IDENTIFIED BY '密码';

//创建哈希值密码的MySQL用户
SELECT password('密码');

//得到密码的哈希值：
*E04600BB4BC876C4A82BB0860EDE6F7829AE5D2A

//设置哈希密码
CREATE USER '用户名'@'localhost' IDENTIFIED BY PASSWORD '*E04600BB4BC876C4A82BB0860EDE6F7829AE5D2A';

//设置哈希密码
CREATE USER '用户名'@'localhost' IDENTIFIED BY PASSWORD '*E04600BB4BC876C4A82BB0860EDE6F7829AE5D2A';

//查看MySQL用户
use mysql
SELECT host,user,password FROM user;

//添加用户名和哈希值密码
CREATE USER '用户名'@'localhost' IDENTIFIED BY PASSWORD '*8633022DF27518A5AE6F1954EE4120D7D99D3524';

//添加DBA授权用户
GRANT ALL ON 数据库名.* TO '用户名'@'localhost';

//重新加载授权表：
FLUSH PRIVILEGES;

//删除空用户
DROP USER 'root'@'localhost';
DROP USER 'root'@'::1';
DROP USER 'root'@'127.0.0.1';
