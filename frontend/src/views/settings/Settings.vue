<template>
  <main id="settings">
    <div id="title">
      <h2>Settings</h2>
    </div>

    <h2>User Profile</h2>
    <profile></profile>

    <h2>Global Users Profile</h2>
    <global></global>
  </main>
</template>

<script>
import { mapState } from 'vuex'
import Profile from './Profile'
import Global from './Global'

export default {
  name: 'settings',
  components: {
    Profile,
    Global
  },
  computed: mapState([ 'user' ])
}
</script>
