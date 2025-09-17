<template>
  <div class="home-container">
    <h1>欢迎来到首页</h1>
    <div class="button-group">
      <button class="action-button">创建产品</button>
      <button class="action-button">提交视频</button>
    </div>
  </div>
</template>

<script setup>
// 目前不需要脚本逻辑
</script>

<style scoped>
.home-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  text-align: center;
}

h1 {
  margin-bottom: 2rem;
}

.button-group {
  display: flex;
  gap: 1rem;
}

.action-button {
  padding: 0.8rem 1.5rem;
  font-size: 1rem;
  cursor: pointer;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f0f0f0;
  transition: background-color 0.3s;
}

.action-button:hover {
  background-color: #e0e0e0;
}
</style>
