<template>
  <div id="app-layout">
    <nav class="sidebar">
      <router-link to="/products">产品管理</router-link>
      <router-link to="/videos">视频管理</router-link>
    </nav>
    <main class="content">
      <router-view />
    </main>
  </div>
</template>

<script setup>
</script>

<style scoped>
#app-layout {
  display: flex;
  width: 100vw; /* 确保占据整个视口宽度 */
  height: 100vh; /* 确保占据整个视口高度 */
}

.sidebar {
  width: 220px;
  flex-shrink: 0; /* 防止侧边栏被压缩 */
  background-color: var(--color-background-soft);
  padding: 20px;
  display: flex;
  flex-direction: column;
  border-right: 1px solid var(--color-border);
}

.sidebar a {
  margin-bottom: 12px;
  text-decoration: none;
  color: var(--color-text);
  font-size: 16px;
  padding: 8px 12px;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.sidebar a:hover {
  background-color: var(--color-background-mute);
}

.sidebar a.router-link-exact-active {
  color: var(--vt-c-white);
  background-color: #42b983; /* 主题绿色 */
  font-weight: 600;
}

.content {
  flex-grow: 1; /* 占据所有剩余空间 */
  padding: 24px;
  overflow-y: auto;
  min-width: 0;
}
</style>
