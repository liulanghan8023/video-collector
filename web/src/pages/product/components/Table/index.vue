<template>
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>名称</th>
        <th>URL</th>
        <th>状态</th>
        <th>创建时间</th>
        <th class="actions-col">操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="product in products" :key="product.ID">
        <td>{{ product.ID }}</td>
        <td>{{ product.name }}</td>
        <td><a :href="product.url" target="_blank">{{ product.url }}</a></td>
        <td>{{ getStatusText(product.status) }}</td>
        <td>{{ new Date(product.created_at).toLocaleString() }}</td>
        <td class="actions-col">
          <button @click="$emit('setStatus', product)" class="primary-btn">状态设置</button>
          <button @click="$emit('edit', product)" class="secondary-btn">修改</button>
          <button @click="$emit('delete', product.ID)" class="danger-btn">删除</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
  defineProps({
    products: {
      type: Array,
      required: true
    }
  });

  defineEmits(['edit', 'delete', 'complete', 'setStatus']);

  const getStatusText = (status) => {
    const statusMap = {
      video_collecting: '视频素材收集中',
      video_collected: '视频素材收集完成',
      video_downloaded: '视频已下载',
      ai_mixed: 'ai混剪完成'
    };
    return statusMap[status] || status;
  };
</script>

<style scoped>
/* 样式与之前页面中的表格样式相同 */
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
  width: auto; /* Adjust width to fit content */
  white-space: nowrap; /* Prevent buttons from wrapping */
  text-align: center;
}

.actions-col button {
  margin: 0 2px; /* Reduce margin */
  padding: 6px 10px; /* Reduce padding */
}

/* Buttons */
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
