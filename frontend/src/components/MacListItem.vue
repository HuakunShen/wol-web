<template>
  <li class="mac-list-item list-group-item mt-2 d-flex justify-content-between">
    <div>
      <p><strong>Name:</strong> {{ mac.name }}</p>
      <p><strong>Mac Address:</strong> {{ mac.mac }}</p>
    </div>
    <div>
      <p><strong>IP:</strong> {{ mac.ip }}</p>
      <p><strong>Port:</strong> {{ mac.port }}</p>
    </div>

    <div>
      <span data-bs-toggle="tooltip" data-bs-placement="left" title="Wake Up">
        <i class="far fa-arrow-alt-circle-up" v-on:click="wake(mac.id)"></i>
      </span>
      <span data-bs-toggle="tooltip" data-bs-placement="left" title="Delete">
        <i class="fas fa-trash-alt" v-on:click="deleteMac({ id: mac.id })"></i>
      </span>
    </div>
  </li>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapActions } from 'vuex';
export default Vue.extend({
  name: 'Mac-List-Item',
  props: ['mac'],
  methods: {
    ...mapActions(['deleteMac']),
    async wake(id: number) {
      const res = await fetch(`/api/wol/${id}`, {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      });

      if (res.status < 400) {
        console.log(res);
      } else {
        console.error('Fail to delete');
        console.error(res);
      }
    },
  },
});
</script>

<style lang="scss" scoped>
.mac-list-item {
  border-radius: 8px;
  padding: 10px 2rem;
  background-color: #fffaf2;
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
.mac-list-item::before {
  content: '';
  height: 100%;
  width: 10px;
  position: absolute;
  background-color: rgb(139, 206, 38);
  left: -0px;
}
</style>
