<template>
  <div v-if="show" class="modal-overlay" @click.self="$emit('cancel')">
    <div class="modal-content">
      <h2>{{ title }}</h2>
      <p>{{ message }}</p>
      <div class="modal-actions">
        <button @click="$emit('cancel')" class="secondary-btn">取消</button>
        <button @click="$emit('confirm')" class="danger-btn">确认删除</button>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  show: Boolean,
  title: { type: String, default: '确认操作' },
  message: { type: String, default: '您确定要执行此操作吗？' }
});
defineEmits(['confirm', 'cancel']);
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
  z-index: 1010; /* Higher z-index */
}

.modal-content {
  background-color: var(--color-background);
  padding: 20px 30px;
  border-radius: 8px;
  width: 100%;
  max-width: 400px;
  text-align: center;
}

.modal-content h2 {
  color: var(--color-heading);
  margin-top: 0;
  margin-bottom: 12px;
}

.modal-content p {
  color: var(--color-text);
  margin-bottom: 24px;
}

.modal-actions {
  display: flex;
  justify-content: center;
  gap: 15px;
}

/* Buttons */
button {
  padding: 8px 20px;
  border: 1px solid var(--color-border);
  background-color: var(--color-background-soft);
  color: var(--color-text);
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s, border-color 0.3s;
}

.secondary-btn:hover:not(:disabled) {
  background-color: var(--color-background-mute);
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
