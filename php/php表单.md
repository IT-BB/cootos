###创建html表单
html表单是使用form标签和多种用于获取输出的元素创建的。如下：

    <form action="script.php" method="post">
    </form>

就php而言，form标签最重要的属性是`actioon`，它指定将把表单数据发送到哪个页面。表单的`method`属性指定如何把数据发送到处理页面。两个选项（`get`和`post`）

实例：

    <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
    <html xmlns="http://www.w3.org/1999/xhtml" lang="zh-CN">
    <head profile="http://gmpg.org/xfn/11">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>落网音乐</title>
    </head>
    <body>
    <form action="form.php" method="post">
    <fieldset><legend>用户注册</legend>
    <p><label>用户名：<input type="text" name="name" /></label></p>
    <p><label>密码：<input type="password" name="password" /></label></p>
    <p><label for="gender">性别：</label><input type="radio" name="gender" value="M" />男
    <input type="radio" name="gender" value="F" />女</p>
    <p><label>年龄：
    <select name="age">
        <option value="0-29">30岁以下</option>
    	<option value="30-60">30-60岁</option>
    	<option value="60+">60岁以上</option>
    </select></label></p>
    <p><input type="submit" name="submit" value="注册" /></p>
    </fieldset>
    </form>
    </body>
    </html>

你也可以为HTML表单标签制定一种字符编码：

    <form accept-charset="utf-8">

处理HTML表单： form.php

    $name=$_POST['name'];
    $password=$_POST['password'];
    $gender=$_POST['gender'];
    $age=$_POST['age'];
    //接收并打印出输入到HTML表单中的信息
    echo $name . '<br/>' . $password . '<br/>' . $gender . '<br/>' . $age;

###验证表单数据
验证表单数据需要使用条件语句以及许多函数、运算符和表达式。一个常用的标准函数是isset(),它用于测试一个变量是否具有值（包括0、FALSE，或者一个空字符串，但不能是NULL）。    
使用isset()函数的一个问题是：空字符串测试为TRUE,这意味着 它不是验证HTML表单中文本输入和文本框的有效方式。要检查用户输入到文本元素中的内容，可以使用empty()函数。它将检查一个变量是否具有空(empty)值：空字符串、0、NULL或FALSE。     
表单验证的第一个目标是确保在表单元素中输入或选择了某些内容。第二个目标是确保提交的数据具有正确的类型（数字、字符串）、正确的格式（如电子邮件）或特定的可接受值（如$gender应该等于1或0）    

    if(!empty($_POST['name'])){
        $name=$_POST['name'];
    }else{
        $name=NULL;
    	echo '<p class="error">请输入您的用户名！</p>';
    }
    
    if(!empty($_POST['password'])){
    	$password=$_POST['password'];
    }else{
    	$password=NULL;
    	echo '<p class="error">请输入您的密码!</p>';
    }
    
    $gender=$_POST['gender'];
    if($gender == 'M') {
        echo '<p><b>Good day,Sir!</b></p>';
    }elseif($gender == 'F') {
        echo '<p><b>Good day,Madam!</b></p>';
    }else{
        echo '<p><b>您还没有选择性别！</b></p>';
    }
    
    if(!empty($_POST['age'])){
    	$age=$_POST['age'];
    }else{
    	$age=NULL;
    	echo '<p class="error">您还没有选择年龄！</p>';
    }
    
    if ( $name && $password && $gender && $age){
        echo '谢谢您的注册!';
    }else{
        echo '您还有选项没有输入';
    }
