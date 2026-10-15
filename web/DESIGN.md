# Kenya Wild: design

Modelled on the attendee's screenshot of a nature-conservation landing page ("Aeline"): calm,
light and natural, with rounded panels on an off-white page, deep green and a lime accent.
Shows Kenyan sightings only (`country.Kenya` in `web/handlers/home.go`).

## Colours (`:root` in `assets/css/app.css`)

| Variable | Hex | Use |
|----------|-----|-----|
| `--paper` | `#ECEEE8` | page background, theme-color |
| `--panel` | `#F5F6F1` | the big rounded section panels |
| `--card` | `#FFFFFF` | sighting cards, pills, ghost buttons |
| `--ink` | `#1C2117` | text |
| `--muted` | `#5E655A` | secondary text (5.4:1 on panel) |
| `--line` | `#E1E4DA` | borders |
| `--forest` | `#4A7A2A` | highlighted words in headings, icons |
| `--forest-deep` | `#3E6B22` | primary buttons (white text), eyebrow text |
| `--lime` | `#D6EA6A` | header button, "Open Data" box |
| `--lime-ink` | `#1E2A10` | text on lime |

## Lettering

- Headings: Outfit 400/500, tight letter-spacing (-0.02em). One word per heading in `<em>`, shown
  in forest green, not italic.
- Body: Inter 400/500/600. Both from Google Fonts in the layout's `<head>`.

## Shape and spacing

- Panels: radius 28px, padding 64px × `--gutter` (clamp 20–64px), 20px apart.
- Cards: radius 20px; buttons 12px; pills fully round.
- Soft shadow only on hover and floating elements: `0 10px 30px rgba(28,33,23,.08)`.

## Components

- **Header**: acacia logo + "Kenya Wild", centred nav to page anchors, lime pill button. Sits over
  the hero on desktop, above it under 860px.
- **Hero**: eyebrow pill, big two-line heading, lead text, green + ghost buttons, a row of round
  animal thumbnails with the observer count, four stat tiles (last one frosted glass). An
  elephant photo fills the right 60%, faded in from the left, with a camera-icon photo credit and
  three frosted floating animal cards (hidden under 1100px).
- **Featured**: text left, three tall photo cards (3:5) with a white icon circle, name and count,
  then a "Seen most in" strip of counties with pin icons.
- **Sightings**: heading and lime "Open Data, Open Eyes." box, then every sighting as a white card:
  4:3 photo with a white animal pill, place as the title, date / count / observer with small
  green icons, and a "View record →" link.
- **Footer**: logo and line, Explore and The data columns, credit.
- Icons are inline Lucide-style SVG strokes (`components.Icon`), never emoji.
- Photos go through `components.Sized`: "medium" on cards, "small" for thumbnails, "large" for
  the hero and tall cards.

## Avoid

- Fake forms or social links: there's no backend for email sign-up and no accounts.
