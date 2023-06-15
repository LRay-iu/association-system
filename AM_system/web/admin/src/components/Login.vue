<template>
  <div class="login-container">
    <h2>登录</h2>
    <form>
      <div class="form-group">
        <label for="username">ID:</label>
        <input type="text" id="username" v-model="username" class="input-field" />
      </div>
      <div class="form-group">
        <label for="password">密码:</label>
        <input type="password" id="password" v-model="password" class="input-field" />
      </div>
      <button @click="login" class="login-button">登录</button>
    </form>
  </div>
</template>

<style>
.login-container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
}

h2 {
  font-size: 24px;
  margin-bottom: 20px;
}

.form-group {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 10px;
}

label {
  width: 80px;
  margin-right: 10px;
}

.input-field {
  flex-grow: 1;
  border: 2px solid #000;
  padding: 5px;
}

button.login-button {
  width: 200px;
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button.login-button:hover {
  background-color: #0056b3;
}
</style>
