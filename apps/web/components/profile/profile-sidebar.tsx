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
  { label: "Bio", icon: MessageSquare },
  { label: "Communities", icon: Building2 },
  { label: "Posted Listings", icon: Home },
  { label: "Bookmarked Listings", icon: Bookmark },
  { label: "My Account", icon: Star },
];

export function ProfileSidebar() {
  return (
    <aside className="rounded-[2rem] border border-orange-100/70 bg-white/90 p-6 shadow-xl shadow-orange-900/5 backdrop-blur-sm lg:sticky lg:top-28">
      <h2 className="px-2 text-3xl font-bold tracking-tight text-gray-900">
        Settings
      </h2>

      <nav className="mt-8 space-y-3" aria-label="Profile settings">
        {profileNavItems.map((item) => {
          const Icon = item.icon;

          return (
            <button
              key={item.label}
              type="button"
              className={`flex w-full items-center gap-4 rounded-2xl px-5 py-4 text-left text-sm font-bold uppercase tracking-[0.16em] transition-all ${
                item.isActive
                  ? "bg-orange-50 text-brand-orange shadow-sm"
                  : "text-gray-400 hover:bg-brand-cream hover:text-gray-700"
              }`}
            >
              <Icon className="h-5 w-5 shrink-0" />
              <span>{item.label}</span>
            </button>
          );
        })}
      </nav>
    </aside>
  );
}
