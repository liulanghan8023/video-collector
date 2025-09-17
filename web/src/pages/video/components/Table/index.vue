<template>
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>URL</th>
        <th>关联产品</th>
        <th>创建时间</th>
        <th class="actions-col">操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="video in videos" :key="video.ID">
        <td>{{ video.ID }}</td>
        <td><a :href="video.url" target="_blank">{{ video.url }}</a></td>
        <td>{{ video.Product ? video.Product.name : 'N/A' }}</td>
        <td>{{ new Date(video.created_at).toLocaleString() }}</td>
        <td class="actions-col">
          <button @click="$emit('edit', video)" class="secondary-btn">修改</button>
          <button @click="$emit('delete', video.ID)" class="danger-btn">删除</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
  defineProps({
    videos: {
      type: Array,
      required: true
    }
  });

  defineEmits(['edit', 'delete']);
</script>

<style scoped>
/* Styles are identical to ProductTable */
table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

th, td {
  border: 1px solid var(--color-border);
  padding: 12px 15px;
  text-align: left;
}

thead {
  background-color: var(--color-background-soft);
}

th {
  font-weight: 600;
}

td a {
  color: #42b983;
  text-decoration: none;
}

td a:hover {
  text-decoration: underline;
}

.actions-col {
  width: 160px;
  text-align: center;
}

.actions-col button {
  margin: 0 5px;
}

button {
  padding: 8px 12px;
  border: 1px solid var(--color-border);
  background-color: var(--color-background-soft);
  color: var(--color-text);
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s, border-color 0.3s;
}

button:hover:not(:disabled) {
  background-color: var(--color-background-mute);
}

.secondary-btn:hover:not(:disabled) {
  border-color: var(--color-border-hover);
}

.danger-btn {
  background-color: #e53935;
  color: var(--vt-c-white);
  border-color: #e53935;
}

.danger-btn:hover:not(:disabled) {
  background-color: #c62828;
}
</style>
