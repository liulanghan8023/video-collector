<template>
  <div v-if="show" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <h2>{{ isEditing ? '修改视频' : '新增视频' }}</h2>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="product">关联产品</label>
          <select id="product" v-model="editableVideo.product_id" required>
            <option disabled value="">请选择一个产品</option>
            <option v-for="product in products" :key="product.ID" :value="product.ID">
              {{ product.name }}
            </option>
          </select>
        </div>
        <div class="form-group">
          <label for="url">URL</label>
          <textarea id="url" v-model="editableVideo.url" required rows="4"></textarea>
        </div>
        <div class="modal-actions">
          <button type="button" @click="$emit('close')" class="secondary-btn">取消</button>
          <button type="submit" class="primary-btn">保存</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  show: Boolean,
  isEditing: Boolean,
  video: Object,
  products: Array // Pass the list of products for the dropdown
});

const emit = defineEmits(['close', 'save']);

const editableVideo = ref({});

watch(() => props.video, (newVal) => {
  editableVideo.value = { ...newVal };
}, { deep: true });

const handleSubmit = () => {
  emit('save', editableVideo.value);
};
</script>

<style scoped>
/* Styles are identical to Product Edit Modal */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: var(--color-background);
  padding: 20px 30px;
  border-radius: 8px;
  width: 100%;
  max-width: 500px;
}

.modal-content h2 {
  color: var(--color-heading);
  margin-top: 0;
  margin-bottom: 24px;
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  color: var(--color-text);
  font-weight: 500;
}

.form-group input, .form-group select, .form-group textarea {
  width: 100%;
  padding: 10px;
  border-radius: 6px;
  border: 1px solid var(--color-border);
  background-color: var(--color-background-soft);
  color: var(--color-text);
  font-size: 15px;
  font-family: inherit; /* Ensure textarea inherits the font */
  resize: vertical; /* Allow vertical resizing */
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 24px;
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

@media (max-width: 768px) {
  .modal-content {
    padding: 20px;
    width: 90%;
    max-width: 90%;
  }
}
</style>
