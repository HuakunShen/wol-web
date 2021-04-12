<template>
  <div id="home">
    <div class="container">
      <form @submit="submit">
        <div class="row">
          <div class="col-lg-6">
            <div class="form-group">
              <input
                v-model="name"
                type="text"
                class="form-control"
                id="mac-name"
                placeholder="Enter Name"
              />
              <small id="emailHelp" class="form-text text-muted">Name</small>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="form-group">
              <input
                v-model="mac"
                type="text"
                class="form-control"
                id="mac-value"
                placeholder="Enter Mac Address"
              />
              <small id="emailHelp" class="form-text text-muted"
                >Mac Address</small
              >
            </div>
          </div>
          <div class="col-lg-6">
            <div class="form-group">
              <input
                v-model="ip"
                type="text"
                class="form-control"
                id="ip"
                placeholder="IP"
              />
              <small id="emailHelp" class="form-text text-muted">IP</small>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="form-group">
              <input
                v-model="port"
                type="number"
                class="form-control"
                id="port"
                placeholder="Port"
              />
              <small id="emailHelp" class="form-text text-muted">Port</small>
            </div>
          </div>
        </div>
        <button type="submit" class="btn btn-primary btn-block mt-1">
          Add
        </button>
        <div
          v-if="lastMsg()"
          class="mt-3 alert"
          v-bind:class="lastMsg().variant"
          role="alert"
        >
          {{ lastMsg().message }}
        </div>
      </form>
      <ComputerList :computers="this.$store.getters.allComputers" />
    </div>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-property-decorator';
import { mapActions, mapGetters } from 'vuex';
import ComputerList from '../components/ComputerList.vue';
export default Vue.extend({
  components: { ComputerList: ComputerList },
  data() {
    return {
      name: '',
      mac: '',
      ip: '255.255.255.255',
      port: '9',
    };
  },
  watch: {
    '$store.state.auth': {
      deep: true,
      handler: function (newValue, oldValue) {
        if (newValue.isAuth === false) {
          this.$router.push({ path: '/login' });
        }
      },
    },
  },
  methods: {
    ...mapGetters(['isAuth', 'lastMsg']),
    ...mapActions(['addComputer']),
    submit(e: Event) {
      e.preventDefault();
      this.addComputer({
        name: this.name,
        mac: this.mac,
        ip: this.ip,
        port: this.port,
      });
    },
  },
});
</script>

<style lang="scss" scoped>
#home {
  .container {
    margin-top: 10px;
    input {
      margin-top: 5px;
      margin-bottom: 5px;
    }
    button {
      width: 100%;
    }
  }
  @media only screen and (min-width: 1600px) {
    .container {
      padding-left: 15rem;
      padding-right: 15rem;
    }
  }

  @media only screen and (max-width: 1600px) {
    .container {
      padding-left: 20%;
      padding-right: 20%;
    }
  }
  @media only screen and (max-width: 1300px) {
    .container {
      padding-left: 15%;
      padding-right: 15%;
    }
  }
  @media only screen and (max-width: 1000px) {
    .container {
      padding-left: 10%;
      padding-right: 10%;
    }
  }
  @media only screen and (max-width: 800px) {
    .container {
      padding-left: 5%;
      padding-right: 5%;
    }
  }
  @media only screen and (max-width: 600px) {
    .container {
      padding-left: 3%;
      padding-right: 3%;
    }
  }
}
</style>
