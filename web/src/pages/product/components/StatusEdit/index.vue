<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h2>设置状态</h2>
      <form @submit.prevent="saveStatus">
        <div class="form-group">
          <label for="status">状态</label>
          <select id="status" v-model="editableProduct.status">
            <option value="video_collecting">视频素材收集中</option>
            <option value="video_collected">视频素材收集完成</option>
            <option value="video_downloaded">视频已下载</option>
            <option value="ai_mixed">ai混剪完成</option>
          </select>
        </div>
        <div class="modal-actions">
          <button type="button" @click="$emit('close')">取消</button>
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
  product: Object,
});

const emit = defineEmits(['close', 'save']);

const editableProduct = ref({});

watch(() => props.product, (newVal) => {
  editableProduct.value = { ...newVal };
});

const saveStatus = () => {
  emit('save', editableProduct.value);
};
</script>

<style scoped>
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
  max-width: 400px;
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

.form-group select {
  width: 100%;
  padding: 10px;
  border-radius: 6px;
  border: 1px solid var(--color-border);
  background-color: var(--color-background-soft);
  color: var(--color-text);
  font-size: 15px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 24px;
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

@media (max-width: 768px) {
  .modal-content {
    padding: 20px;
    width: 90%;
    max-width: 90%;
  }
}
</style>
