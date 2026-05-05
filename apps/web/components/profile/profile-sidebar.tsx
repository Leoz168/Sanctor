import {
  Bookmark,
  Building2,
  Home,
  MessageSquare,
  Star,
  User,
} from "lucide-react";

const profileNavItems = [
  { label: "Profile", icon: User, isActive: true },
  { label: "My Listings", icon: Home },
  { label: "My Communities", icon: Building2 },
  { label: "Bookmark", icon: Bookmark },
  { label: "My Account", icon: Star },
];

export function ProfileSidebar() {
  return (
    <aside className="rounded-[2rem] border border-orange-100 bg-white/95 p-5 shadow-xl shadow-orange-900/8 backdrop-blur-sm lg:sticky lg:top-24">
      <h2 className="px-2 text-3xl font-bold tracking-tight text-gray-900">
        Settings
      </h2>

      <nav className="mt-6 space-y-2" aria-label="Profile settings">
        {profileNavItems.map((item) => {
          const Icon = item.icon;

          return (
            <button
              key={item.label}
              type="button"
              className={`flex w-full items-center gap-3 rounded-2xl px-4 py-3 text-left text-xs font-bold uppercase tracking-[0.14em] transition-all ${
                item.isActive
                  ? "bg-brand-cream text-brand-orange shadow-sm shadow-orange-900/5"
                  : "text-gray-500 hover:bg-brand-cream hover:text-gray-700"
              }`}
            >
              <Icon className="h-4 w-4 shrink-0" />
              <span>{item.label}</span>
            </button>
          );
        })}
      </nav>
    </aside>
  );
}
