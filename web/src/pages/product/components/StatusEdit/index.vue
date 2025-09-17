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
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: var(--color-background);
  padding: 20px;
  border-radius: 8px;
  width: 400px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-group select {
  width: 100%;
  padding: 8px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>
