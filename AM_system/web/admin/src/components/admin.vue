<template>
    <div id="app">
        <div id="sidebar" :class="{ 'sidebar-hidden': isSidebarHidden }">
            <h3>目录</h3>
            <ul>
                <li><a href="#clubs">社团</a></li>
                <li><a href="#activities">社团活动</a></li>
                <li><a href="#members">社团成员</a></li>
            </ul>
            <button class="logout-button" @click="logout">退出登录</button>
        </div>

        <div id="content">
            <h2 id="clubs">社团</h2>
            <button class="publish-button" @click="publishActivity">发布活动</button>
            <button class="deleteass">注销社团</button>
            <table>
                <thead>
                    <tr>
                        <th>社团ID</th>
                        <th>社团名称</th>
                        <th>社团信息</th>
                        <th>社团人数</th>
                    </tr>
                </thead>
                <tbody v-if="clubs">
                    <tr>
                        <td>{{ clubs.AssID }}</td>
                        <td>{{ clubs.AssName }}</td>
                        <td>{{ clubs.AssNote }}</td>
                        <td>{{ clubs.AssScale }}</td>
                    </tr>
                </tbody>
            </table>

            <h2 id="activities">社团活动</h2>
            <table>
                <thead>
                    <tr>
                        <th>活动编号</th>
                        <th>活动名称</th>
                        <th>负责人</th>
                        <th>日期</th>
                        <th>活动内容</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="activity in activities" :key="activity.id">
                        <td>{{ activity.id }}</td>
                        <td>{{ activity.name }}</td>
                        <td>{{ activity.personInCharge }}</td>
                        <td>{{ activity.date }}</td>
                        <td>{{ activity.content }}</td>
                        <td>
                            <button class="delete-button" @click="deleteActivity(activity.id)">删除</button>
                        </td>
                    </tr>
                </tbody>
            </table>

            <h2>社团成员</h2>
            <table>
                <thead>
                    <tr>
                        <th>学号</th>
                        <th>学生姓名</th>
                        <th>性别</th>
                        <th>学院</th>
                        <th>专业</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="member in members" :key="member.studentId">
                        <td>{{ member.studentId }}</td>
                        <td>{{ member.studentName }}</td>
                        <td>{{ member.gender }}</td>
                        <td>{{ member.college }}</td>
                        <td>{{ member.major }}</td>
                        <td>
                            <button class="delete-button" @click="deleteMember(member.studentId)">删除</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
  
<script>
import axios from 'axios';

export default {
    data() {
        return {
            clubs: null,
            activities: [
                {
                    id: 1,
                    name: '太极拳培训',
                    personInCharge: '张三',
                    date: '2023-06-10',
                    content: '太极拳是一种中国传统的武术，注重以柔克刚、以静制动的原理。本次培训将由资深太极拳教练亲自指导，欢迎对太极拳感兴趣的同学参加。'
                },
                {
                    id: 2,
                    name: '武术表演',
                    personInCharge: '李四',
                    date: '2023-06-15',
                    content: '武术表演将集结各类武术门派的精彩演出，展示中国武术的博大精深和独特魅力。届时会有精彩的拳法、剑术、器械等表演节目，欢迎大家观看。'
                },
                // {
                //     id: 3,
                //     name: '看老奶奶跳舞',
                //     personInCharge: '蔡晓坤',
                //     date: '2023-06-9',
                //     content: '看老奶奶跳舞，学习如何养生'
                // },
            ],
            members: [
                // {
                //     studentId: '20210001',
                //     studentName: '王五',
                //     gender: '男',
                //     college: '文学院',
                //     major: '汉语言文学'
                // },
                {
                    studentId: '20210002',
                    studentName: '赵六',
                    gender: '女',
                    college: '理学院',
                    major: '数学'
                },
                {
                    studentId: '20210003',
                    studentName: '刘七',
                    gender: '男',
                    college: '工学院',
                    major: '计算机科学与技术'
                },
                {
                    studentId: '20210014',
                    studentName: '张十八',
                    gender: '女',
                    college: '人文学院',
                    major: '历史学'
                },
                {
                    studentId: '20210015',
                    studentName: '李十九',
                    gender: '男',
                    college: '法学院',
                    major: '法学'
                },
            ]
        };
    },
    mounted() {
        this.fetchData();
    },
    methods: {
        fetchData() {
            axios.get('http://localhost:8080/api/v1/ass/search?AssName=武术')
                .then(response => {
                    this.clubs = response.data.data;
                    // this.activities = response.data.activities;
                    // this.members = response.data.members;
                })
                .catch(error => {
                    console.error(error);
                });
        },
        toggleSidebar() {
            this.isSidebarHidden = !this.isSidebarHidden;
        },
        joinClub() {
            const url = 'http://localhost:8080/api/v1/ass/search';
            const data = {
                AssID: this.content
            };
            axios.get(url, data)
                .then(response => {
                    // 查询成功后的逻辑处理
                    // 可以更新页面数据或进行其他操作
                    console.log('查询成功');
                })
                .catch(error => {
                    // 处理查询失败的情况
                    console.error(error);
                });
        },
        logout() {
            // 处理退出登录的逻辑
        },
        deleteClub(clubId) {
            const url = `http://localhost:8080/api/v1/ass/${clubId}`;
            axios.delete(url)
                .then(response => {
                    // 删除成功后的逻辑处理
                    // 可以更新页面数据或进行其他操作
                    console.log('删除成功');
                })
                .catch(error => {
                    // 处理删除失败的情况
                    console.error(error);
                });
        },
        deleteActivity(activityId) {
            // 处理删除活动的逻辑
        },
        deleteMember(studentId) {
            // 处理删除成员的逻辑
        },
        publishActivity() {
            // 处理发布活动的逻辑
        }
    }
};
</script>
  
<style>
body {
    background-color: #e0e0e0;
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
}

#app {
    display: flex;
    height: 100vh;
    overflow: hidden;
}

#sidebar {
    background-color: #f2f2f2;
    width: 200px;
    padding: 20px;
    transition: margin-left 0.3s;
    z-index: 1;
}

ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
}

li {
    margin-bottom: 10px;
}

#content {
    padding: 20px;
    flex-grow: 1;
    overflow-y: auto;
}

.logout-button {
    position: absolute;
    top: 20px;
    right: 20px;
    padding: 10px;
    background-color: #f2f2f2;
    border: none;
    cursor: pointer;
}

.sidebar-hidden {
    margin-left: -200px;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
}

th,
td {
    padding: 8px;
    text-align: left;
    border-bottom: 1px solid #ddd;
}

th {
    background-color: #c9c4db;
    font-weight: bold;
}

#search-bar {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
}

#search-bar input[type="text"] {
    flex-grow: 1;
    padding: 5px;
    border: 1px solid #ddd;
    border-radius: 5px;
    margin-right: 5px;
}

#search-bar i {
    color: #777;
    font-size: 16px;
    cursor: pointer;
}

.delete-button {
    padding: 5px 10px;
    background-color: #ff0000;
    border: none;
    color: #fff;
    cursor: pointer;
}

.publish-button {
    padding: 5px 10px;
    background-color: #007bff;
    border: none;
    color: #fff;
    cursor: pointer;
    border-radius: 5px;
    margin-left: 10px;
}

.deleteass {
    padding: 5px 10px;
    background-color: #ff4a03;
    border: none;
    color: #fff;
    cursor: pointer;
    border-radius: 5px;
    margin-left: 10px;
}
</style>
  