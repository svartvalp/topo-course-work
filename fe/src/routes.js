import Booklist from './components/Booklist'
import Cart from './components/Cart'

const routes = [
    {
        path: "/",
        component: Booklist
    },
    {
        path: "/cart",
        component: Cart
    }
]

export default routes