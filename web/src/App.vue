<template>
  <div id="app-layout">
    <!-- Hamburger Menu Button (Mobile Only) -->
    <button @click="toggleSidebar" class="hamburger-menu">
      <span></span>
      <span></span>
      <span></span>
    </button>

    <!-- Sidebar -->
    <nav class="sidebar" :class="{ 'is-open': isSidebarOpen }">
      <router-link to="/products" @click="isSidebarOpen = false">产品管理</router-link>
      <router-link to="/videos" @click="isSidebarOpen = false">视频管理</router-link>
    </nav>

    <!-- Main Content -->
    <main class="content">
      <router-view />
    </main>

    <!-- Backdrop -->
    <div v-if="isSidebarOpen" @click="toggleSidebar" class="backdrop"></div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const isSidebarOpen = ref(false);

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value;
};
</script>

<style scoped>
#app-layout {
  display: flex;
  width: 100vw;
  height: 100vh;
}

.sidebar {
  width: 220px;
  flex-shrink: 0;
  background-color: var(--color-background-soft);
  padding: 20px;
  display: flex;
  flex-direction: column;
  border-right: 1px solid var(--color-border);
  transition: transform 0.3s ease-in-out;
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
  background-color: #42b983;
  font-weight: 600;
}

.content {
  flex-grow: 1;
  padding: 0;
  overflow-y: auto;
  min-width: 0;
}

.hamburger-menu {
  display: none; /* Hidden by default */
}

.backdrop {
  display: none; /* Hidden by default */
}

/* Responsive Layout: Mobile */
@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    height: 100%;
    transform: translateX(-100%);
    z-index: 1000;
    border-right: 1px solid var(--color-border);
  }

  .sidebar.is-open {
    transform: translateX(0);
  }

  .hamburger-menu {
    display: block;
    position: fixed;
    top: 15px;
    left: 15px;
    z-index: 1001;
    background: var(--color-background-soft);
    border: 1px solid var(--color-border);
    border-radius: 6px;
    padding: 8px;
    cursor: pointer;
  }

  .hamburger-menu span {
    display: block;
    width: 22px;
    height: 3px;
    background-color: var(--color-text);
    margin: 4px 0;
    border-radius: 1px;
  }

  .content {
    padding: 60px 15px 15px 15px; /* Adjust padding to account for hamburger menu */
  }

  .backdrop {
    display: block;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    z-index: 999;
  }
}
</style>
