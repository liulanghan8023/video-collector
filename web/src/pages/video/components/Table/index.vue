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
        <td data-label="ID">{{ video.ID }}</td>
        <td data-label="URL"><a :href="video.url" target="_blank" class="url-link">{{ video.url }}</a></td>
        <td data-label="关联产品">{{ video.Product ? video.Product.name : 'N/A' }}</td>
        <td data-label="创建时间">{{ new Date(video.created_at).toLocaleString() }}</td>
        <td data-label="操作" class="actions-col">
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
  max-width: 250px; /* Adjust this width as needed */
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle; /* Helps align the link nicely within the cell */
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

/* Mobile Responsive Styles */
@media (max-width: 768px) {
  .url-link {
    max-width: 180px; /* Allow a bit more width on mobile cards */
  }

  thead {
    display: none; /* Hide table headers on mobile */
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
    text-align: right; /* Align content to the right */
    padding-left: 50%; /* Create space for the label */
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
    width: calc(50% - 30px);
    padding-right: 10px;
    white-space: nowrap;
    text-align: left;
    font-weight: 600;
    color: var(--color-heading);
  }

  .actions-col {
    width: 100%;
    text-align: right; /* Align buttons to the right */
  }

  .actions-col button {
    margin: 5px 0 5px 5px;
  }
}
</style>
