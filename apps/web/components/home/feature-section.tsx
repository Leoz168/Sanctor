import { Mail, MessageSquare, ShoppingBag } from "lucide-react";
import { FeatureCard } from "@/components/home/feature-card";

export function FeatureSection() {
  return (
    <section className="max-w-7xl mx-auto px-4 pb-24">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8" id="feature-columns">
        <FeatureCard
          icon={Mail}
          title="Subscribe"
          description="Get notified as soon as new property listings that match your criteria become available."
        >
          <div className="relative">
            <input
              type="email"
              placeholder="your@email.com"
              className="w-full bg-gray-50 border border-gray-100 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-brand-orange/20 mb-3"
            />
            <button className="w-full py-3 bg-gray-900 text-white rounded-xl text-sm font-bold hover:bg-black transition-colors">
              Notify Me
            </button>
          </div>
        </FeatureCard>

        <FeatureCard
          icon={MessageSquare}
          title="Forums"
          description="Connect with other students, find roommates, and discuss neighborhood safety or landlords."
        >
          <ul className="space-y-3">
            {["Roommate Search", "General Housing Discussion", "Landlord Reviews"].map((topic) => (
              <li
                key={topic}
                className="flex items-center gap-2 text-sm font-medium text-gray-500 hover:text-brand-orange cursor-pointer"
              >
                <div className="w-1.5 h-1.5 rounded-full bg-brand-orange/30" />
                {topic}
              </li>
            ))}
          </ul>
        </FeatureCard>

        <FeatureCard
          icon={ShoppingBag}
          title="Market"
          description="Buy and sell second-hand furniture, textbooks, and essentials specifically for student living."
        >
          <div className="grid grid-cols-2 gap-3">
            {["Furniture", "Books"].map((label) => (
              <div
                key={label}
                className="h-20 bg-gray-50 rounded-xl flex items-center justify-center border border-dashed border-gray-200"
              >
                <span className="text-[10px] uppercase tracking-widest font-bold text-gray-400">
                  {label}
                </span>
              </div>
            ))}
          </div>
        </FeatureCard>
      </div>
    </section>
  );
}
