<template>
  <main id="users">
    <div id="title">
      <h2>{{ $t('settings.users') }}</h2>

      <div id="actions">
        <router-link class="action" to="/users/new" :aria-label="$t('buttons.new')" :title="$t('buttons.new')">
          <i class="material-icons">person_add</i>
          <span>{{ $t('buttons.new') }}</span>
        </router-link>
      </div>
    </div>

    <div class="content full">
      <table>
        <tr>
          <th></th>
          <th>{{ $t('settings.username') }}</th>
          <th>{{ $t('settings.scope') }}</th>
          <th></th>
        </tr>

        <tr v-for="user in users" :key="user.id">
          <td><i v-if="user.perm.admin" class="material-icons" style="color: #d44">verified_user</i><i v-else class="material-icons">account_circle</i></td>
          <td>{{ user.username }}</td>
          <td>{{ user.scope }}</td>
          <td class="small">
            <router-link :to="'/users/' + user.id"><i class="material-icons">mode_edit</i></router-link>
          </td>
        </tr>
      </table>
    </div>
  </main>
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
