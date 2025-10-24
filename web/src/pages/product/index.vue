<template>
  <div>
    <div class="header">
      <div class="actions">
        <div class="search-container">
          <select v-model="selectedStatus" @change="fetchProducts" class="status-filter">
            <option value="">所有状态</option>
            <option value="video_collecting">视频素材收集中</option>
            <option value="video_collected">视频素材收集完成</option>
            <option value="video_downloaded">视频已下载</option>
            <option value="ai_mixed">ai混剪完成</option>
          </select>
        </div>
        <button @click="openModal()" class="primary-btn">新增产品</button>
      </div>
    </div>

    <Table
      :products="products"
      @edit="openModal"
      @delete="promptDelete"
      @setStatus="handleSetStatus"
      @add-video="openAddVideoModal"
    />

    <div class="pagination">
      <button @click="prevPage" :disabled="page <= 1">上一页</button>
      <span>第 {{ page }} 页 / 共 {{ totalPages }} 页</span>
      <button @click="nextPage" :disabled="page >= totalPages">下一页</button>
    </div>

    <Edit
      :show="showModal"
      :is-editing="isEditing"
      :product="currentProduct"
      @close="closeModal"
      @save="handleSave"
    />

    <ConfirmModal
      :show="showConfirmModal"
      title="删除产品"
      message="您确定要永久删除此产品吗？此操作无法撤销。"
      @cancel="cancelDelete"
      @confirm="executeDelete"
    />

    <StatusEdit
      :show="showStatusModal"
      :product="productForStatusUpdate"
      @close="showStatusModal = false"
      @save="handleSaveStatus"
    />

    <AddVideo
      :show="showAddVideoModal"
      :product="productForVideo"
      @close="closeAddVideoModal"
      @save="handleSaveVideo"
    />

  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import Table from './components/Table/index.vue';
import Edit from './components/Edit/index.vue';
import StatusEdit from './components/StatusEdit/index.vue';
import AddVideo from './components/AddVideo/index.vue';
import ConfirmModal from '../../components/ConfirmModal/index.vue'; // Import shared component

const products = ref([]);
const page = ref(1);
const pageSize = ref(10);
const total = ref(0);
const selectedStatus = ref('');

// Edit/Create Modal state
const showModal = ref(false);
const isEditing = ref(false);
const currentProduct = ref({});

// Add Video Modal state
const showAddVideoModal = ref(false);
const productForVideo = ref({});

// Status Edit Modal state
const showStatusModal = ref(false);
const productForStatusUpdate = ref({});

// Confirm Delete Modal state
const showConfirmModal = ref(false);
const productIdToDelete = ref(null);

const totalPages = computed(() => Math.ceil(total.value / pageSize.value));

const fetchProducts = async () => {
  try {
    let url = `/api/products?page=${page.value}&pageSize=${pageSize.value}`;
    if (selectedStatus.value) {
      url += `&status=${selectedStatus.value}`;
    }
    const response = await fetch(url);
    if (!response.ok) throw new Error('Network response was not ok');
    const result = await response.json();
    products.value = result.data;
    total.value = result.total;
  } catch (error) {
    console.error('Failed to fetch products:', error);
  }
};

const openModal = (product = null) => {
  if (product && product.ID) {
    isEditing.value = true;
    currentProduct.value = { ...product };
  } else {
    isEditing.value = false;
    currentProduct.value = { name: '', url: '' };
  }
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const handleSave = async (productToSave) => {
  const method = isEditing.value ? 'PUT' : 'POST';
  const url = isEditing.value ? `/api/products/${productToSave.ID}` : '/api/products';

  try {
    const response = await fetch(url, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(productToSave),
    });
    if (!response.ok) throw new Error('Failed to save product');
    closeModal();
    await fetchProducts();
  } catch (error) {
    console.error('Error saving product:', error);
  }
};

const handleSetStatus = (product) => {
  productForStatusUpdate.value = product;
  showStatusModal.value = true;
};

const handleSaveStatus = async (productToUpdate) => {
  try {
    const response = await fetch(`/api/products/${productToUpdate.ID}/status`,
        {
          method: 'PUT',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ status: productToUpdate.status }),
        });
    if (!response.ok) throw new Error('Failed to update product status');
    showStatusModal.value = false;
    await fetchProducts();
  } catch (error) {
    console.error('Error updating product status:', error);
  }
};

const openAddVideoModal = (product) => {
  productForVideo.value = product;
  showAddVideoModal.value = true;
};

const closeAddVideoModal = () => {
  showAddVideoModal.value = false;
};

const handleSaveVideo = async ({ productId, urls }) => {
  try {
    const response = await fetch(`/api/products/${productId}/videos`,
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ urls }),
        });
    if (!response.ok) throw new Error('Failed to save videos');
    closeAddVideoModal();
  } catch (error) {
    console.error('Error saving videos:', error);
  }
};



// --- Delete Logic ---
const promptDelete = (id) => {
  productIdToDelete.value = id;
  showConfirmModal.value = true;
};

const cancelDelete = () => {
  showConfirmModal.value = false;
  productIdToDelete.value = null;
};

const executeDelete = async () => {
  if (productIdToDelete.value === null) return;

  try {
    const response = await fetch(`/api/products/${productIdToDelete.value}`, { method: 'DELETE' });
    if (!response.ok) throw new Error('Failed to delete product');
    await fetchProducts();
  } catch (error) {
    console.error('Error deleting product:', error);
  }
  cancelDelete(); // Close modal and reset ID
};
// --- End Delete Logic ---

const prevPage = () => {
  if (page.value > 1) {
    page.value--;
    fetchProducts();
  }
};

const nextPage = () => {
  if (page.value < totalPages.value) {
    page.value++;
    fetchProducts();
  }
};

onMounted(() => {
  fetchProducts();
});
</script>

<style scoped>
/* Global styles for the page */
h1 {
  color: var(--color-heading);
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  flex-wrap: wrap;
  gap: 15px; /* Add gap for spacing when wrapped */
}

.actions {
  display: flex;
  flex-grow: 1;
  justify-content: flex-start;
  gap: 10px;
}

.search-container {
  display: flex;
  gap: 10px;
}

.status-filter {
  padding: 8px 12px;
  border: 1px solid var(--color-border);
  background-color: var(--color-background-soft);
  color: var(--color-text);
  border-radius: 6px;
}

.pagination {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 20px; /* Added margin top */
}

/* Button styles are now mostly in components, but we can keep some base styles */
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