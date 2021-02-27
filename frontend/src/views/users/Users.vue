<template>
  <article id="users">
    <header id="title">
      <h1>{{ $t('sidebar.users') }}</h1>

      <router-link class="action" to="/users/new" :aria-label="$t('buttons.new')" :title="$t('buttons.new')">
        <i class="material-icons">person_add</i>
      </router-link>
    </header>

    <main>
      <router-link v-for="user in users" :key="user.id" :to="'/users/' + user.id" class="div">
        <i v-if="user.perm.admin" class="material-icons" style="color: #d44">verified_user</i><i v-else class="material-icons">account_circle</i>
        <div class="name">{{ user.username }}</div>
        <div class="scope">{{ user.scope }}</div>
      </router-link>
    </main>
  </article>
</template>

<script>
import { users as api } from '@/api'

export default {
  name: 'users',
  data: function () {
    return {
      users: []
    }
  },
  async created () {
    try {
      this.users = await api.getAll()
    } catch (e) {
      this.$showError(e)
    }
  }
}
</script>
