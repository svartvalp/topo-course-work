import Vue from "vue"
import Vuex from "vuex"
import axios from "axios";

Vue.use(Vuex)

export const getters = {
    books: state => {
        return state.books
    },
    session: state => {
        return state.session
    },
    cart: state => {
        return state.cart
    }
}


export const mutations = {
    updateBooks(state, payload) {
        state.books = payload
    },
    updateSession(state, session) {
        state.session = session
    },
    updateCart(state, payload) {
        state.cart = payload
    }
}

export const actions = {
    loadBooks(context, session) {
        console.log(`loading books, session ${session}`)
        axios.get(`/api/books?page=1&size=100&session=${session}`).then((response) => {
            context.commit('updateBooks', response.data)
        })
    },
    async loadSession(context) {
        console.log('loading session')
        let session = localStorage.getItem('session')
        if (!session) {
            let response = await axios.post('/api/session/generate')
            session = response.data.session
            localStorage.setItem('session', session)
        }
        context.commit('updateSession', session)
    },
    loadCart(context, session) {
        axios.get(`/api/cart?session=${session}`).then((response) => {
            context.commit('updateCart', response.data)
        })
    },
    addToCart(context, payload) {
        axios.post(`/api/cart/add?session=${payload.session}&bookID=${payload.id}`).then(() => {
            axios.get(`/api/cart?session=${payload.session}`).then((response) => {
                context.commit('updateCart', response.data)
            })
            axios.get(`/api/books?page=1&size=100&session=${payload.session}`).then((response) => {
                context.commit('updateBooks', response.data)
            })
        })
    },
    removeFromCart(context, payload) {
        axios.post(`/api/cart/remove?session=${payload.session}&bookID=${payload.id}`).then(() => {
            axios.get(`/api/cart?session=${payload.session}`).then((response) => {
                context.commit('updateCart', response.data)
            })
            axios.get(`/api/books?page=1&size=100&session=${payload.session}`).then((response) => {
                context.commit('updateBooks', response.data)
            })
        })
    },
    clearCart(context, session) {
        axios.post(`/api/cart/clear?session=${session}`).then(() => {
            axios.get(`/api/cart?session=${session}`).then((response) => {
                context.commit('updateCart', response.data)
            })
            axios.get(`/api/books?page=1&size=100&session=${session}`).then((response) => {
                context.commit('updateBooks', response.data)
            })
        })
    }
}



export default new Vuex.Store({
    state: {
        books: [],
        session: null,
        cart: null
    },
    getters,
    mutations,
    actions
})