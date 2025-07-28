import React, { useEffect, useState } from 'react'
import axios from 'axios'

export default function App() {
  const [products, setProducts] = useState([])
  const [inventory, setInventory] = useState([])
  const [name, setName] = useState("")
  const [selectedProduct, setSelectedProduct] = useState("")
  const [quantity, setQuantity] = useState(0)
  const [type, setType] = useState("IN")

  const fetchData = async () => {
    try{
    const [pRes, iRes] = await Promise.all([
      axios.get(`${import.meta.env.VITE_API_URL}/products`),
      axios.get(`${import.meta.env.VITE_API_URL}/inventory`)
    ])
    setProducts(pRes.data)
    setInventory(iRes.data)
  } catch (err) {
    console.error("Error loading data", err)
  }
  useEffect(() => {
    fetchData()
  }, [])

  const addProduct = async () => {
    if (!name) return
    await axios.post(`${import.meta.env.VITE_API_URL}/products`, { name })
    setName("")
    fetchData()
  }

  const recordTransaction = async () => {
    if (!selectedProduct || !quantity) return
    await axios.post(`${import.meta.env.VITE_API_URL}/transaction`, {
      transactionType: type,
      transactionDate: new Date().toISOString(),
      items: [{ productId: parseInt(selectedProduct), quantity: parseInt(quantity) }]
    })
    setSelectedProduct("")
    setQuantity(0)
    fetchData()
  }

  return (
    <div className="p-8 max-w-4xl mx-auto">
      <h1 className="text-2xl font-bold mb-6">Warehouse Inventory Tracker</h1>

      <div className="mb-8 space-y-4">
        <h2 className="text-xl font-semibold">Add Product</h2>
        <div className="flex gap-2">
          <input
            className="border p-2 flex-1"
            value={name}
            placeholder="Product name"
            onChange={(e) => setName(e.target.value)}
          />
          <button className="bg-blue-600 text-white px-4 py-2 rounded" onClick={addProduct}>Add</button>
        </div>
      </div>

      <div className="mb-8 space-y-4">
        <h2 className="text-xl font-semibold">Record Transaction</h2>
        <div className="flex flex-wrap gap-2 items-center">
          <select className="border p-2" value={selectedProduct} onChange={e => setSelectedProduct(e.target.value)}>
            <option value="">Select Product</option>
            {products.map(p => <option key={p.id} value={p.id}>{p.name}</option>)}
          </select>
          <input
            type="number"
            className="border p-2 w-24"
            placeholder="Qty"
            value={quantity}
            onChange={e => setQuantity(e.target.value)}
          />
          <select className="border p-2" value={type} onChange={e => setType(e.target.value)}>
            <option value="IN">IN</option>
            <option value="OUT">OUT</option>
          </select>
          <button className="bg-green-600 text-white px-4 py-2 rounded" onClick={recordTransaction}>Submit</button>
        </div>
      </div>

      <div>
        <h2 className="text-xl font-semibold mb-4">Current Inventory</h2>
        <table className="min-w-full bg-white border">
          <thead>
            <tr className="bg-gray-100">
              <th className="py-2 px-4 border-b">Product</th>
              <th className="py-2 px-4 border-b">Quantity</th>
            </tr>
          </thead>
          <tbody>
            {inventory && inventory.length > 0 ? (
              inventory.map((inv) => (
              <tr key={inv.productID}>
                <td className="py-2 px-4 border-b">{inv.name}</td>
                <td className="py-2 px-4 border-b">{inv.quantity}</td>
              </tr>
            ))
          ) : (
            !loading && <p>No inventory available.</p>
          )}
          </tbody>
        </table>
      </div>
    </div>
  )
}
