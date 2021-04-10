<template>
  <li
    class="computer-list-item list-group-item mt-2 d-flex justify-content-between"
  >
    <div>
      <p><strong>Name:</strong> {{ computer.name }}</p>
      <p><strong>Mac Address:</strong> {{ computer.mac }}</p>
    </div>
    <div>
      <p><strong>IP:</strong> {{ computer.ip }}</p>
      <p><strong>Port:</strong> {{ computer.port }}</p>
    </div>

    <div>
      <span data-bs-toggle="tooltip" data-bs-placement="left" title="Wake Up">
        <i
          class="far fa-arrow-alt-circle-up"
          v-on:click="wake(computer.id)"
        ></i>
      </span>
      <span data-bs-toggle="tooltip" data-bs-placement="left" title="Delete">
        <i
          class="fas fa-trash-alt"
          v-on:click="deleteComputer({ id: computer.id })"
        ></i>
      </span>
    </div>
  </li>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapActions } from 'vuex';
export default Vue.extend({
  name: 'Computer-List-Item',
  props: ['computer'],
  methods: {
    ...mapActions(['deleteComputer']),
    async wake(id: number) {
      const res = await fetch(`/api/wol/${id}`, {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      });

      if (res.status < 400) {
        console.log('success');
      } else {
        console.error('Fail to delete');
        console.error(res);
      }
    },
  },
});
</script>

<style lang="scss" scoped>
.computer-list-item {
  border-radius: 8px;
  padding: 10px 2rem;
  background-color: #fffaf7;
  display: flex;
  align-items: center;
  justify-content: space-between;
  .fas.fa-trash-alt {
    cursor: pointer;
    transform: translateY(5px);
    color: red;
    padding-left: 1rem;
    font-size: 1.3rem;
  }
  .far.fa-arrow-alt-circle-up {
    cursor: pointer;
    transform: translateY(5px);
    font-size: 1.5rem;
  }
}
.computer-list-item::before {
  content: '';
  height: 100%;
  width: 10px;
  position: absolute;
  background-color: rgb(139, 206, 38);
  left: -0px;
}
</style>
