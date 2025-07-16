import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { BACKEND_URL } from '../config';
import { Input, TextField } from '@mui/material';

export default function ProductTable() {
    const [data, setData] = React.useState([])
    const [error, setError] = React.useState<undefined | string>()
    const [loading, setLoading] = React.useState(false)
    const [search, setSearch] = React.useState<undefined | string>()

    const fetchData = async () => {
        try {
            setLoading(true)
            const url = `${BACKEND_URL}/products`;
            const res = await fetch(url)

            setData(await res.json())
        } catch (e) {
            if (typeof e === "string") {
                setError(e.toUpperCase())
            } else if (e instanceof Error) {
                setError(e.message)
            }
        }
        finally {
            setLoading(false)
        }
    }

    React.useEffect(() => {
        fetchData();
    }, []);

    React.useEffect(() => {

        const searchProducts = async () => {
            if (!search || search.trim().length === 0) {
                fetchData();
                return
            }
            try {
                setLoading(true)
                const res = await fetch(`${BACKEND_URL}/products/search?query=${search}`)

                setData(await res.json())
            } catch (e) {
                if (typeof e === "string") {
                    setError(e.toUpperCase())
                } else if (e instanceof Error) {
                    setError(e.message)
                }
            }
            finally {
                setLoading(false)
            }
        }

        searchProducts();
    }, [search])


    if (error) {
        return <div>{error}</div>
    }

    return (
        <Paper sx={{ m: 5 }}>
            <TextField variant="outlined" fullWidth placeholder="Search" value={search} onChange={(e) => setSearch(e.target.value)} />
            <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell align="right">Description</TableCell>
                            <TableCell align="right">Category</TableCell>
                            <TableCell align="right">Brand</TableCell>
                            <TableCell align="right">Stock Quantity</TableCell>
                            <TableCell align="right">SKU</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {data.map((row) => (
                            <TableRow
                                key={row["id"]}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">
                                    {row["name"]}
                                </TableCell>
                                <TableCell align="right">{row["description"]}</TableCell>
                                <TableCell align="right">{row["category"]}</TableCell>
                                <TableCell align="right">{row["brand"]}</TableCell>
                                <TableCell align="right">{row["stock_quantity"]}</TableCell>
                                <TableCell align="right">{row["sku"]}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Paper>

    );
}