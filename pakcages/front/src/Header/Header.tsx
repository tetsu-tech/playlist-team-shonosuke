import './Header.css';
import { HeaderRight } from "./HeaderRight";

export function Header() {
  return (
    <div className="Header">
      <div>What's now</div>
      <HeaderRight />
    </div>
  );
}
