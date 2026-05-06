import { Bookmark, Building2, Home, Star, User } from "lucide-react";

export type ProfileTab = "profile" | "listings" | "communities" | "bookmarks" | "account";

const profileNavItems: Array<{
  id: ProfileTab;
  label: string;
  icon: typeof User;
  count?: number;
}> = [
  { id: "profile", label: "Profile", icon: User },
  { id: "listings", label: "My Listings", icon: Home, count: 1 },
  { id: "communities", label: "My Communities", icon: Building2, count: 2 },
  { id: "bookmarks", label: "Bookmark", icon: Bookmark },
  { id: "account", label: "My Account", icon: Star },
];

interface ProfileSidebarProps {
  activeTab: ProfileTab;
  onTabChange: (tab: ProfileTab) => void;
}

export function ProfileSidebar({
  activeTab,
  onTabChange,
}: ProfileSidebarProps) {
  return (
    <aside className="h-[720px] rounded-[2rem] border border-orange-100 bg-white/95 p-5 shadow-xl shadow-orange-900/8 backdrop-blur-sm lg:sticky lg:top-24">
      <h2 className="px-2 text-3xl font-bold tracking-tight text-gray-900">
        Settings
      </h2>

      <nav className="mt-6 space-y-2" aria-label="Profile settings">
        {profileNavItems.map((item) => {
          const Icon = item.icon;
          const isActive = item.id === activeTab;

          return (
            <button
              key={item.id}
              type="button"
              onClick={() => onTabChange(item.id)}
              className={`flex w-full items-center gap-3 rounded-2xl px-4 py-3 text-left text-xs font-bold uppercase tracking-[0.14em] transition-all ${
                isActive
                  ? "bg-brand-cream text-brand-orange shadow-sm shadow-orange-900/5"
                  : "text-gray-500 hover:bg-brand-cream hover:text-gray-700"
              }`}
            >
              <Icon className="h-4 w-4 shrink-0" />
              <span className="min-w-0 flex-1">{item.label}</span>
              {item.count ? (
                <span className="flex h-6 min-w-6 items-center justify-center rounded-full bg-brand-orange px-2 text-xs font-bold text-white">
                  {item.count}
                </span>
              ) : null}
            </button>
          );
        })}
      </nav>
    </aside>
  );
}
