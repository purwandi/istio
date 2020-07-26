import express from 'express'
import axios from 'axios'

// create new app
const app = express()
const port = 8080

// Set engine configuration
app.use(express.static('public'))
app.use(express.json())
app.use(express.urlencoded({ extended: true }))
app.set('view engine', 'pug')

// Register the route
app.get('/', async (req, res) => {
  try {
    const products = await axios.get(`${process.env.PRODUCT_SERVICE}/product`)
    return res.render('index', { products: products.data })
  } catch (error) {
    return res.json({ error: error.toString() })
  }
})

app.get('/:id', async(req, res) => {
  try {
    const product = await axios.get(`${process.env.PRODUCT_SERVICE}/product/${req.params.id}`)
    const reviews = await axios.get(`${process.env.REVIEW_SERVICE}/review/${req.params.id}`)

    return res.render('detail', { product: product.data, reviews: reviews.data })
  } catch (error) {
    return res.json({ error: error.toString() })
  }
})

app.post("/review/:id", async(req, res) => {
  try {
    const data = {
      comment: req.body.comment,
      rating: req.body.rating,
      product_id: req.params.id
    }
    await axios.post(`${process.env.REVIEW_SERVICE}/review/${req.params.id}`, data)

    return res.redirect(`/${req.params.id}`)
  } catch (error) {
    return res.json({ error: error.toString() })
  }
})

// Start the server
app.listen(port, ()=> console.log('App started'))