<template>
    <div class="register-container">
        <h2>发布活动</h2>
        <form>
            <div class="form-group">
                <label for="student-id">活动ID:</label>
                <input type="number" id="student-id" v-model="studentId" class="input-field" />
            </div>
            <div class="form-group">
                <label for="student-name">活动名称:</label>
                <input type="text" id="student-name" v-model="studentName" class="input-field" />
            </div>
            <div class="form-group">
                <label for="student-sex">负责人:</label>
                <input type="text" id="student-sex" v-model="sex" class="input-field" />
            </div>
            <div class="form-group">
                <label for="college">活动时间:</label>
                <input type="text" id="college" v-model="college" class="input-field" />
            </div>
            <div class="form-group2">
                <label for="major">活动内容:</label>
                <input type="text" id="major" v-model="major" class="input-field2" />
            </div>
            <button @click="checkPasswords" class="register-button">登记并发布</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    data() {
        return {
            studentId: null,
            studentName: '',
            sex: '',
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
                "Sex": this.sex,
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
    justify-content: top;
    margin-bottom: 10px;
}

.form-group2 {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 30px;
}

.input-field2 {
    flex-grow: 1;
    border: 2px solid #000;
    padding: 5px;
    height: 100px;
    /* 调整文本框的高度 */
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
}</style>