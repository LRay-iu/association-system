<template>
  <div id="app">
    <div id="sidebar" :class="{ 'sidebar-hidden': isSidebarHidden }">
      <h3>目录</h3>
      <ul>
        <li><a href="#clubs">社团</a></li>
        <li><a href="#activities">社团活动</a></li>
      </ul>
      <button class="logout-button" @click="logout">退出登录</button>
    </div>

    <div id="content">
      <h2 id="clubs">社团</h2>
      <div id="search-bar">
        <input type="text" placeholder="搜索社团">
        <i class="fas fa-search"></i>
        <button class="join-button" @click="joinClub">加入</button>

      </div>
      <table>
        <thead>
          <tr>
            <th>社团ID</th>
            <th>社团名称</th>
            <th>社团信息</th>
            <th>社团人数</th>
            <th>操作</th> <!-- 新增的列 -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="club in clubs" :key="club.AssID">
            <td>{{ club.AssID }}</td>
            <td>{{ club.AssName }}</td>
            <td>{{ club.AssNote }}</td>
            <td>{{ club.AssScale }}</td>
            <td>
              <button class="delete-button" @click="deleteClub(club.AssID)">退出社团</button>
            </td>
          </tr>
        </tbody>
      </table>

      <h2 id="activities">社团活动</h2>
      <table>
        <thead>
          <tr>
            <th>活动名称</th>
            <th>负责人</th>
            <th>日期</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="activity in activities" :key="activity.id">
            <td>{{ activity.name }}</td>
            <td>{{ activity.personInCharge }}</td>
            <td>{{ activity.date }}</td>
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
      clubs: [],
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
      ],
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      axios.get('http://localhost:8080/api/v1/ass/ask')
        .then(response => {
          this.clubs = response.data.data;
          // this.activities = response.data.activities;
        })
        .catch(error => {
          console.error(error);
        });
    },
    toggleSidebar() {
      this.isSidebarHidden = !this.isSidebarHidden;
    },
    joinClub() {
      const url = `http://localhost:8080/api/v1/ass/search`;
      const data = {
        "AssID": this.content
      }
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

.extra-content {
  margin-top: 40px;
  padding: 20px;
  background-color: #f9f9f9;
}

.extra-content h2 {
  color: #333;
  font-size: 24px;
}

.extra-content p {
  color: #555;
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
</style>
  