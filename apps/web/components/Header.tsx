const links = [
  { href: "#communities", label: "Communities" },
  { href: "#messaging", label: "Messaging" },
  { href: "#discovery", label: "Discovery" },
];

export function Header() {
  return (
    <header className="site-header">
      <div className="brand">
        <div className="brand-mark">S</div>
        <div className="brand-copy">
          <span className="brand-title">Sanctor</span>
          <span className="brand-subtitle">Campus communities, rebuilt.</span>
        </div>
      </div>
      <nav className="nav-links" aria-label="Primary">
        {links.map((link) => (
          <a className="nav-link" href={link.href} key={link.label}>
            {link.label}
          </a>
        ))}
        <a className="nav-pill" href={process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}>
          API
        </a>
      </nav>
    </header>
  );
}
