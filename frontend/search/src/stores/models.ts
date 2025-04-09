export interface Response {
  result: Result;
}

export interface Result {
  searchRequestId: string
  text: string
  misspellCorrected: boolean
  lastPage: boolean
  total: number
  perPage: number
  results: SearchResult[]
  filters: Filter[]
  responseType: string
}

export interface SearchResult {
  type: string
  track: Track
}

export interface Track {
  id: string
  realId: string
  title: string
  version?: string
  major: Major
  available: boolean
  availableForPremiumUsers: boolean
  availableFullWithoutPermission: boolean
  availableForOptions: string[]
  disclaimers: string[]
  storageDir: string
  durationMs: number
  fileSize: number
  r128: R128
  fade: Fade
  previewDurationMs: number
  artists: Artist[]
  albums: Album[]
  coverUri: string
  derivedColors: DerivedColors
  ogImage: string
  lyricsAvailable: boolean
  type: string
  rememberPosition: boolean
  trackSharingFlag: string
  lyricsInfo: LyricsInfo
  trackSource: string
  specialAudioResources: string[]
  contentWarning?: string
  chart?: Chart
}

export interface Major {
  id: number
  name: string
}

export interface R128 {
  i: number
  tp: number
}

export interface Fade {
  inStart: number
  inStop: number
  outStart: number
  outStop: number
}

export interface Artist {
  id: number
  name: string
  various: boolean
  composer: boolean
  available: boolean
  cover: Cover
  genres: string[]
  disclaimers: string[]
}

export interface Cover {
  type: string
  uri: string
  prefix: string
}

export interface Album {
  id: number
  title: string
  type?: string
  metaType: string
  version?: string
  year: number
  releaseDate: string
  coverUri: string
  ogImage: string
  genre: string
  trackCount: number
  likesCount?: number
  recent: boolean
  veryImportant: boolean
  artists: Artist[]
  labels: Label[]
  available: boolean
  availableForPremiumUsers: boolean
  availableForOptions: string[]
  availableForMobile: boolean
  availablePartially: boolean
  bests: number[]
  disclaimers: string[]
  listeningFinished: boolean
  trackPosition: TrackPosition
  contentWarning?: string
}

export interface Label {
  id: number
  name: string
}

export interface TrackPosition {
  volume: number
  index: number
}

export interface DerivedColors {
  average: string
  waveText: string
  miniPlayer: string
  accent: string
}

export interface LyricsInfo {
  hasAvailableSyncLyrics: boolean
  hasAvailableTextLyrics: boolean
}

export interface Chart {
  position: number
  progress: string
  listeners: number
  shift: number
}
export interface Filter {
  id: string
  displayName: string
}
