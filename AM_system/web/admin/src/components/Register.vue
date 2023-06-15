<template>
  <div class="register-container">
    <h2>学生注册</h2>
    <form>
      <div class="form-group">
        <label for="student-id">学生ID:</label>
        <input type="number" id="student-id" v-model="studentId" class="input-field" />
      </div>
      <div class="form-group">
        <label for="student-name">学生姓名:</label>
        <input type="text" id="student-name" v-model="studentName" class="input-field" />
      </div>
      <div class="form-group">
        <label for="student-sex">性别:</label>
        <input type="text" id="student-sex" v-model="sex" class="input-field" />
      </div>
      <div class="form-group">
        <label for="college">所属学院:</label>
        <input type="text" id="college" v-model="college" class="input-field" />
      </div>
      <div class="form-group">
        <label for="major">所属专业:</label>
        <input type="text" id="major" v-model="major" class="input-field" />
      </div>
      <div class="form-group">
        <label for="password">密码:</label>
        <input type="password" id="password" v-model="password" class="input-field" />
      </div>
      <div class="form-group">
        <label for="confirm-password">确认密码:</label>
        <input type="password" id="confirm-password" v-model="confirmPassword" class="input-field" />
      </div>
      <button @click="checkPasswords" class="register-button">注册</button>
    </form>
    <div v-if="passwordMismatch" class="error-message">两次密码不一致</div>
    <div v-if="registrationSuccess" class="success-message">注册成功</div>
  </div>
</template>

<script>
import axios from 'axios';  
export default {
  data() {
    return {
      studentId: null,
      studentName: '',
      sex:'',
      college: '',
      major: '',
      password: '',
      confirmPassword: '',
      passwordMismatch: false,
      registrationSuccess: false,
    };
  },
  methods: {
    checkPasswords() {
      if (this.password !== this.confirmPassword) {
        this.passwordMismatch = true;
        this.registrationSuccess = false;
      } else {
        this.passwordMismatch = false;
        this.registrationSuccess = true;
        this.sendRegistrationData();
      }
    },
    sendRegistrationData() {
      const url = 'http://localhost:8080/api/v1/stu/add';
      const data = {
        "StuId": this.studentId,
        "StuName": this.studentName,
        "Sex":this.sex,
        "College": this.college,
        "Major": this.major,
        "StuPassword": this.password
      };

      axios.post(url, data)
      console.log(data)
        .then(response => {
          // 处理成功响应
          console.log('注册成功');
        })
        .catch(error => {
          // 处理错误
          console.error('注册失败:', error);
        });
    }
  },
};

</script>

<style>
.register-container {
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
  width: 120px;
  margin-right: 10px;
}

.input-field {
  flex-grow: 1;
  border: 2px solid #000;
  padding: 5px;
}

.error-message {
  color: red;
  margin-top: 5px;
}

.success-message {
  color: green;
  margin-top: 5px;
}

button.register-button {
  width: 200px;
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button.register-button:hover {
  background-color: #0056b3;
}
</style>
