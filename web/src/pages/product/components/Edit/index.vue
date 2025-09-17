<template>
  <div v-if="show" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <h2>{{ isEditing ? '修改产品' : '新增产品' }}</h2>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="name">名称</label>
          <input type="text" id="name" v-model="editableProduct.name" required>
        </div>
        <div class="form-group">
          <label for="url">URL</label>
          <input type="text" id="url" v-model="editableProduct.url" required>
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
  product: Object
});

const emit = defineEmits(['close', 'save']);

const editableProduct = ref({});

watch(() => props.product, (newVal) => {
  editableProduct.value = { ...newVal };
}, { deep: true });

const handleSubmit = () => {
  emit('save', editableProduct.value);
};
</script>

<style scoped>
/* Modal styles are self-contained */
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

.form-group input {
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

.secondary-btn:hover:not(:disabled) {
  border-color: var(--color-border-hover);
}
</style>
