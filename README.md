用戶端  

V 1. 使用者可以瀏覽目前的留言 GET /user/comment  

V 2. 使用者可以註冊帳號 POST /user/signup  

V 3. 使用者可以登入 POST /user/authentication  

V 4. 使用者登入後可以留言 POST /user/comment  

V 5. 使用者登入後回覆特定留言，但只開放針對留言做回覆，不能回覆一則回覆  POST /user/comment/13/reply  

管理介面 is_superuser = true  

V 1. 管理者可以看到目前的留言並搜尋留言內容 GET /superuser/comment?search=${string}

V 2. 管理者可以隱藏留言 PUT /superuser/comment  

V 3. 管理者可以將使用者停權 PUT /superuser/user  

V 4. 管理員可以鎖定文章不可回覆 PUT /superuser/comment

![Screenshot from 2022-04-18 08-50-01](https://user-images.githubusercontent.com/37808888/163738802-0fccd82a-8e73-42ed-9dbb-b03dd308a901.png)
