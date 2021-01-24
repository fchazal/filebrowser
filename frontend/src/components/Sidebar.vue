<template>
  <nav :class="{active}">
    <template v-if="isLogged">
      <router-link class="action" to="/files/" :aria-label="$t('sidebar.home')" :title="$t('sidebar.home')">
        <i class="material-icons">cloud</i>
        <span>{{ $t('sidebar.home') }}</span>
      </router-link>

      <div id="hierarchy">

      </div>

      <div v-for="(drive, index) in this.user.drives" :key="index">
        <router-link class="action" :to="'/files/' + drive.name" :aria-label="drive.name" :title="drive.name">
          <i class="material-icons">{{ drive.icon }}</i>
          <span>{{ drive.name }}</span>
        </router-link>
      </div>

      <div class="separator"></div>

      <router-link class="action" to="/shares" :aria-label="$t('sidebar.shares')" :title="$t('sidebar.shares')">
        <i class="material-icons">folder_shared</i>
        <span>{{$t('sidebar.shares')}}</span>
      </router-link>

      <router-link v-if="user.perm.admin" class="action" to="/users" :aria-label="$t('sidebar.users')" :title="$t('sidebar.users')">
        <i class="material-icons">people</i>
        <span>{{$t('sidebar.users')}}</span>
      </router-link>
    </template>

    <template v-else>
      <router-link class="action" to="/login" :aria-label="$t('sidebar.login')" :title="$t('sidebar.login')">
        <i class="material-icons">exit_to_app</i>
        <span>{{ $t('sidebar.login') }}</span>
      </router-link>

      <router-link v-if="signup" class="action" to="/login" :aria-label="$t('sidebar.signup')" :title="$t('sidebar.signup')">
        <i class="material-icons">person_add</i>
        <span>{{ $t('sidebar.signup') }}</span>
      </router-link>
    </template>
  </nav>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import { version, signup, disableExternal, noAuth, authMethod } from '@/utils/constants'

export default {
  name: 'sidebar',
  computed: {
    ...mapState([ 'user' ]),
    ...mapGetters([ 'isLogged' ]),
    active () {
      return this.$store.state.show === 'sidebar'
    },
    signup: () => signup,
    version: () => version,
    disableExternal: () => disableExternal,
    noAuth: () => noAuth,
    authMethod: () => authMethod
  },
  methods: {
  }
}
</script>
