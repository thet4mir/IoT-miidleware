import Keycloak from 'keycloak-js'

const kc = new Keycloak({
  realm: process.env.REACT_APP_IAM_REALM,
  clientId: process.env.REACT_APP_IAM_CLIENT_ID,
  url: process.env.REACT_APP_IAM_ENDPOINT,
  'enable-cors': true
})

const initKeycloak = (onAuthenticatedCallback) => {
  kc.init({
    onLoad: 'login-required',
  })
  .then((authenticated) => {
    onAuthenticatedCallback()
  })
  .catch(console.error)
}

const doLogout = () => kc.logout({ redirectUri: process.env.REACT_APP_ENV === 'dev' ? 'http://localhost:3000' : 'http://mapper.andorean.com' })

const UserService = {
  initKeycloak,
  doLogout,
}

export default UserService
