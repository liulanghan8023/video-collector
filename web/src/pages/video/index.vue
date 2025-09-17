<template>
  <div>
    <div class="header flex-wrap">
      <h1>视频管理</h1>
      <button @click="openModal()" class="primary-btn">新增视频</button>
    </div>

    <div class="filters">
      <div class="filter-group">
        <label for="product-search">按产品筛选</label>
        <select id="product-search" v-model="searchProductId">
          <option value="">所有产品</option>
          <option v-for="product in productList" :key="product.ID" :value="product.ID">
            {{ product.name }}
          </option>
        </select>
      </div>
    </div>

    <Table
      :videos="videos"
      @edit="openModal"
      @delete="promptDelete"
    />

    <div class="pagination">
      <button @click="prevPage" :disabled="page <= 1">上一页</button>
      <span>第 {{ page }} 页 / 共 {{ totalPages }} 页</span>
      <button @click="nextPage" :disabled="page >= totalPages">下一页</button>
    </div>

    <Edit
      :show="showModal"
      :is-editing="isEditing"
      :video="currentVideo"
      :products="productList"
      @close="closeModal"
      @save="handleSave"
    />

    <ConfirmModal
      :show="showConfirmModal"
      title="删除视频"
      message="您确定要永久删除此视频吗？此操作无法撤销。"
      @cancel="cancelDelete"
      @confirm="executeDelete"
    />

  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import Table from './components/Table/index.vue';
import Edit from './components/Edit/index.vue';
import ConfirmModal from '../../components/ConfirmModal/index.vue';

// List state
const videos = ref([]);
const page = ref(1);
const pageSize = ref(10);
const total = ref(0);
const searchProductId = ref('');

// Data for dropdowns
const productList = ref([]);

// Modal states
const showModal = ref(false);
const isEditing = ref(false);
const currentVideo = ref({});
const showConfirmModal = ref(false);
const videoIdToDelete = ref(null);

const totalPages = computed(() => Math.ceil(total.value / pageSize.value));

// Watch for search filter changes
watch(searchProductId, () => {
  page.value = 1; // Reset to first page when filter changes
  fetchVideos();
});

const fetchVideos = async () => {
  try {
    let url = `/api/videos?page=${page.value}&pageSize=${pageSize.value}`;
    if (searchProductId.value) {
      url += `&product_id=${searchProductId.value}`;
    }
    const response = await fetch(url);
    if (!response.ok) throw new Error('Network response was not ok');
    const result = await response.json();
    videos.value = result.data;
    total.value = result.total;
  } catch (error) {
    console.error('Failed to fetch videos:', error);
  }
};

const fetchAllProducts = async () => {
  try {
    const response = await fetch(`/api/products`);
    if (!response.ok) throw new Error('Network response was not ok');
    productList.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch all products:', error);
  }
};

const openModal = (video = null) => {
  if (video && video.ID) {
    isEditing.value = true;
    currentVideo.value = { ...video };
  } else {
    isEditing.value = false;
    currentVideo.value = { url: '', product_id: '' };
  }
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const handleSave = async (videoToSave) => {
  const method = isEditing.value ? 'PUT' : 'POST';
  const url = isEditing.value ? `/api/videos/${videoToSave.ID}` : '/api/videos';

  try {
    const response = await fetch(url, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(videoToSave),
    });
    if (!response.ok) throw new Error('Failed to save video');
    closeModal();
    await fetchVideos();
  } catch (error) {
    console.error('Error saving video:', error);
  }
};

const promptDelete = (id) => {
  videoIdToDelete.value = id;
  showConfirmModal.value = true;
};

const cancelDelete = () => {
  showConfirmModal.value = false;
  videoIdToDelete.value = null;
};

const executeDelete = async () => {
  if (videoIdToDelete.value === null) return;
  try {
    const response = await fetch(`/api/videos/${videoIdToDelete.value}`, { method: 'DELETE' });
    if (!response.ok) throw new Error('Failed to delete video');
    await fetchVideos();
  } catch (error) {
    console.error('Error deleting video:', error);
  }
  cancelDelete();
};

const prevPage = () => {
  if (page.value > 1) {
    page.value--;
    fetchVideos();
  }
};

const nextPage = () => {
  if (page.value < totalPages.value) {
    page.value++;
    fetchVideos();
  }
};

onMounted(() => {
  fetchVideos();
  fetchAllProducts();
});
</script>

<style scoped>
/* Using a more specific selector for h1 to avoid affecting component h2 */
:deep(h1) {
  color: var(--color-heading);
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.filters {
  margin-bottom: 20px;
  padding: 15px;
  background-color: var(--color-background-soft);
  border-radius: 8px;
}

.filter-group label {
  margin-right: 10px;
  font-weight: 500;
}

.filter-group select {
  padding: 8px;
  border-radius: 6px;
  border: 1px solid var(--color-border);
  background-color: var(--color-background);
}

.pagination {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 20px;
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

button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.primary-btn {
  background-color: #42b983;
  color: var(--vt-c-white);
  border-color: #42b983;
}

.primary-btn:hover:not(:disabled) {
  background-color: #36a471;
}
</style>