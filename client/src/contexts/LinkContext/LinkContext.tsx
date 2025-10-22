import { createContext, useContext, useState, type ReactNode } from "react"

type LinkContextType = {
    originalUrl: string
    setOriginalUrl: (url: string) => void
    shortenedUrl: string
    setShortenedUrl: (url: string) => void
    loading: boolean
    setLoading: (value: boolean) => void
    error: string
    setError: (msg: string) => void
}


const LinkContext = createContext<LinkContextType | undefined>(undefined)

export function LinkProvider({ children }: { children: ReactNode }) {
  const [originalUrl, setOriginalUrl] = useState('')
  const [shortenedUrl, setShortenedUrl] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  return (
    <LinkContext.Provider
      value={{
        originalUrl,
        setOriginalUrl,
        shortenedUrl,
        setShortenedUrl,
        loading,
        setLoading,
        error,
        setError,
      }}
    >
      {children}
    </LinkContext.Provider>
  )
}

export function useLink() {
  const context = useContext(LinkContext)
  if (!context) throw new Error('useUrl deve ser usado dentro de UrlProvider')
  return context
}