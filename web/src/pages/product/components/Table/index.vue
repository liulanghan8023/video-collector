<template>
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>名称</th>
        <th>URL</th>
        <th>视频数</th>
        <th>状态</th>
        <th>创建时间</th>
        <th class="actions-col">操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="product in products" :key="product.ID">
        <td data-label="ID">{{ product.ID }}</td>
        <td data-label="名称">{{ product.name }}</td>
        <td data-label="URL"><a :href="product.url" target="_blank" class="url-link">{{ product.url }}</a></td>
        <td data-label="视频数">{{ product.video_count }}</td>
        <td data-label="状态">{{ getStatusText(product.status) }}</td>
        <td data-label="创建时间">{{ new Date(product.created_at).toLocaleString() }}</td>
        <td data-label="操作" class="actions-col">
          <button @click="$emit('add-video', product)" class="primary-btn">添加视频</button>
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
/* Desktop Styles */
table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

th, td {
  border: 1px solid var(--color-border);
  padding: 12px 15px;
  text-align: left;
  vertical-align: middle;
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

.url-link {
  display: inline-block;
  max-width: 250px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle;
}

.actions-col {
  width: auto;
  white-space: nowrap;
  text-align: center;
}

.actions-col button {
  margin: 0 2px;
  padding: 6px 10px;
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

.primary-btn {
    background-color: #42b983;
    color: var(--vt-c-white);
    border-color: #42b983;
}

.primary-btn:hover:not(:disabled) {
    background-color: #36a471;
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

/* Mobile Responsive Styles */
@media (max-width: 768px) {
  .url-link {
    max-width: 150px;
  }

  thead {
    display: none;
  }

  tr {
    display: block;
    margin-bottom: 15px;
    border: 1px solid var(--color-border);
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  }

  td {
    display: block;
    text-align: right;
    padding-left: 40%;
    position: relative;
    border: none;
    border-bottom: 1px solid var(--color-border-hover);
  }

  td:last-child {
    border-bottom: none;
  }

  td::before {
    content: attr(data-label);
    position: absolute;
    left: 15px;
    width: calc(40% - 30px);
    padding-right: 10px;
    white-space: nowrap;
    text-align: left;
    font-weight: 600;
    color: var(--color-heading);
  }

  .actions-col {
    width: 100%;
    text-align: right;
  }

  .actions-col button {
    margin: 5px 0 5px 5px;
  }
}
</style>
