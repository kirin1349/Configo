import { Num3 } from "./Num3";

export class Num4 extends Num3
{
    w: number = 0;

    constructor(x: number = 0, y: number = 0, z: number = 0, w: number = 0)
    {
        super(x, y, z);

        this.w = w;
    }
}