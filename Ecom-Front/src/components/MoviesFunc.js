import React, { useEffect, useState, Fragment } from 'react'
import { Link } from 'react-router-dom'

function MoviesFunc(props) {
  const [movies, setMovies] = useState([])
  const [error, setError] = useState('')

  useEffect(() => {
    console.log(process.env.REACT_APP_API_URL)
    // fetch(`${process.env.REACT_APP_API_URL}/v1/movies`)
    fetch(`${process.env.REACT_APP_API_URL}/v1/movies`)
      .then((response) => {
        console.log('Status code is', response.status)
        if (response.status !== 200) {
          //   let err = Error
          //   err.message = 'Invalid response code ' + response.status
          //   this.setState({ error: err })
          setError('Invalid response code: ', response.status)
        } else {
          setError(null)
        }
        return response.json()
      })
      .then((json) => {
        // this.setState(
        //   {
        //     movies: json.movies,
        //     isLoaded: true,
        //   },
        //   (error) => {
        //     this.setState({
        //       isLoaded: true,
        //       error,
        //     })
        //   }
        // )
        setMovies(json.movies)
      })
  }, []) // *** must set default value
  if (error !== null) {
    return <div>Error: {error.message}</div>
  } else {
    return (
      <Fragment>
        <h2>Choose a movie</h2>
        <div className='list-group'>
          {movies.map((m) => (
            <Link
              key={m.id}
              className='list-group-item list-group-item-action'
              to={`/movies/${m.id}`}
            >
              {m.title}
            </Link>
          ))}
        </div>
      </Fragment>
    )
  }
}

export default MoviesFunc
