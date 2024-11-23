import { Router } from "vue-router";
declare module "*.svg" {
  import type React from "react";

  const content: React.FC<React.SVGProps<SVGElement>>;
  export default content;
}

declare const process: {
  env: {
    VITE_API_ROOT: string;
  };
};
declare module "*.png" {
  const value: string;
  export default value;
}
declare module "*.webp" {
  const value: string;
  export default value;
}
declare module "*.jpg" {
  const value: string;
  export default value;
}

declare module "vue" {
  interface ComponentCustomProperties {
    $router: Router;
  }
}
